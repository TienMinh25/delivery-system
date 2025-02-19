package thirdparty

import (
	"context"
	"os"
	"strings"

	"github.com/TienMinh25/delivery-system/pkg"
	minio "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/pkg/errors"
)

type storage struct {
	client     *minio.Client
	bucketName string
	urlPrefix  string
	region     string
}

// initialize minio
func NewStorage() (pkg.Storage, error) {
	endpoint := os.Getenv("MINIO_ENDPOINT_URL")
	accessKeyID := os.Getenv("MINIO_ACCESS_KEY")
	secretAccessKey := os.Getenv("MINIO_SECRET_KEY")
	bucketName := os.Getenv("MINIO_BUCKET_AVATARS")
	bucketPolicy := os.Getenv("MINIO_BUCKET_AVATARS_POLICY")
	region := os.Getenv("MINIO_REGION")
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
		Region: region,
	})

	if err != nil {
		return nil, errors.Wrap(err, "minio.New")
	}

	ctx := context.Background()

	isBucketExists, err := minioClient.BucketExists(ctx, bucketName)

	if err == nil && !isBucketExists {
		err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{
			Region: region,
		})

		if err != nil {
			return nil, errors.Wrap(err, "minioClient.MakeBucket")
		}

		err = minioClient.SetBucketPolicy(ctx, bucketName, bucketPolicy)

		if err != nil {
			return nil, errors.Wrap(err, "minioClient.SetBucketPolicy")
		}
	}

	endpointURL := minioClient.EndpointURL()
	urlPrefix := endpointURL.Scheme + "://" + bucketName + ".s3." + region + "." + endpointURL.Host + "/"

	return &storage{
		client:     minioClient,
		bucketName: bucketName,
		urlPrefix:  urlPrefix,
		region:     region,
	}, nil
}

// Delete implements pkg.Storage.
func (s *storage) Delete(ctx context.Context, name string) error {
	// get key object (get path of file on s3)
	name = strings.TrimPrefix(name, s.urlPrefix)

	err := s.client.RemoveObject(ctx, s.bucketName, name, minio.RemoveObjectOptions{})

	if err != nil {
		return errors.Wrap(err, "s.client.RemoveObject")
	}

	return nil
}

// Upload implements pkg.Storage.
func (s *storage) Upload(ctx context.Context, payload pkg.UploadInput) (string, error) {
	info, err := s.client.PutObject(ctx, s.bucketName, payload.Name, payload.File, payload.Size, minio.PutObjectOptions{
		ContentType: payload.ContentType,
	})

	// if err exists, upload to minio fail
	if err != nil {
		return "", errors.Wrap(err, "s.client.PutObject")
	}

	return s.urlPrefix + info.Key, nil
}
