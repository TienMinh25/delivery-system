package gateway

import (
	"bytes"
	"context"
	"image"
	"image/jpeg"
	"image/png"
	"io"

	"github.com/TienMinh25/delivery-system/internal/products"
	"github.com/TienMinh25/delivery-system/pkg"
	"github.com/disintegration/imaging"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/pkg/errors"
)

const _MAX_AVATAR_BYTES = 5 << 20

type Service interface {
	// service handles for user management
	signUp(ctx context.Context, payload *signUpRequest) (*signUpResponse, error)
	confirmSignUp(ctx context.Context, payload *confirmSignUpRequest) error
	signIn(ctx context.Context, payload *signInRequest) (*token, error)
	checkToken(ctx context.Context, accessToken string) (*pkg.User, error)
	refreshToken(ctx context.Context, payload *refreshTokenRequest) (*token, error)
	updateUser(ctx context.Context, user *user, payload *updateUserRequest) error
	confirmChangeEmail(ctx context.Context, payload *confirmChangeEmailRequest) error
	updateAvatar(ctx context.Context, user *user, file pkg.File, fileInfo pkg.FileInfo) error
	forgotPassword(ctx context.Context, userID string) error
	resetPassword(ctx context.Context, payload *resetPasswordRequest) error
	changePassword(ctx context.Context, user *user, payload *changePasswordRequest) error
	signOut(ctx context.Context, userID string, refreshTokenPayload *refreshTokenRequest) error
	deleteUser(context.Context, string) error

	// service handles for other services
	getPartnerProducts(ctx context.Context) (*products.GetAllResponse, error)
	checkOrder(ctx context.Context, user *user, payload *checkRequest) error
	confirmOrder(ctx context.Context, user *user, payload *confirmRequest) (*confirmResponse, error)
	payOrder(ctx context.Context, payload *payRequest) (*payResponse, error)
	pickUpOrder(ctx context.Context, payload *pickupRequest) error
	assignOrder(ctx context.Context, user *user, payload *assignRequest) error
}

type service struct {
	authManager pkg.Auth
	storage     pkg.Storage
}

func NewService(
	authManager pkg.Auth,
) Service {
	return &service{
		authManager: authManager,
	}
}

// assignOrder implements Service.
func (s *service) assignOrder(ctx context.Context, user *user, payload *assignRequest) error {
	panic("unimplemented")
}

// changePassword implements Service.
func (s *service) changePassword(ctx context.Context, user *user, payload *changePasswordRequest) error {
	panic("unimplemented")
}

// checkOrder implements Service.
func (s *service) checkOrder(ctx context.Context, user *user, payload *checkRequest) error {
	panic("unimplemented")
}

// checkToken implements Service.
func (s *service) checkToken(ctx context.Context, accessToken string) (*pkg.User, error) {
	user, err := s.authManager.CheckToken(ctx, accessToken)

	if err != nil {
		return nil, errors.Wrap(err, "s.authManager.CheckToken")
	}

	return user, nil
}

// confirmChangeEmail implements Service.
func (s *service) confirmChangeEmail(ctx context.Context, payload *confirmChangeEmailRequest) error {
	err := s.authManager.ConfirmChangeEmail(ctx, payload.AccessToken, &pkg.VerifyEmailRequest{
		UserID: payload.UserID,
		Code:   payload.Code,
	})

	if err != nil {
		return errors.Wrap(err, "s.authManager.ConfirmChangeEmail")
	}

	return nil
}

// confirmOrder implements Service.
func (s *service) confirmOrder(ctx context.Context, user *user, payload *confirmRequest) (*confirmResponse, error) {
	panic("unimplemented")
}

// confirmSignUp implements Service.
func (s *service) confirmSignUp(ctx context.Context, payload *confirmSignUpRequest) error {
	err := s.authManager.ConfirmSignUp(ctx, &pkg.VerifyEmailRequest{
		UserID: payload.UserID,
		Code:   payload.Code,
	})

	if err != nil {
		return errors.Wrap(err, "s.authManager.ConfirmSignUp")
	}

	return nil
}

// deleteUser implements Service.
func (s *service) deleteUser(context.Context, string) error {
	panic("unimplemented")
}

// forgotPassword implements Service.
func (s *service) forgotPassword(ctx context.Context, userID string) error {
	panic("unimplemented")
}

// getPartnerProducts implements Service.
func (s *service) getPartnerProducts(ctx context.Context) (*products.GetAllResponse, error) {
	panic("unimplemented")
}

// payOrder implements Service.
func (s *service) payOrder(ctx context.Context, payload *payRequest) (*payResponse, error) {
	panic("unimplemented")
}

// pickUpOrder implements Service.
func (s *service) pickUpOrder(ctx context.Context, payload *pickupRequest) error {
	panic("unimplemented")
}

// refreshToken implements Service.
func (s *service) refreshToken(ctx context.Context, payload *refreshTokenRequest) (*token, error) {
	authToken, err := s.authManager.RefreshToken(ctx, payload.RefreshToken)

	if err != nil {
		return nil, errors.Wrap(err, "s.authManager.RefreshToken")
	}

	return &token{
		AccessToken:  authToken.AccessToken,
		RefreshToken: authToken.RefreshToken,
		ExpiresIn:    authToken.ExpiresIn,
	}, nil
}

// resetPassword implements Service.
func (s *service) resetPassword(ctx context.Context, payload *resetPasswordRequest) error {
	panic("unimplemented")
}

// signIn implements Service.
func (s *service) signIn(ctx context.Context, payload *signInRequest) (*token, error) {
	_, authToken, err := s.authManager.SignIn(ctx, &pkg.SignInRequest{
		Email:    payload.Email,
		Password: payload.Password,
	})

	if err != nil {
		return nil, errors.Wrap(err, "s.authManager.SignIn")
	}

	return &token{
		AccessToken:  authToken.AccessToken,
		RefreshToken: authToken.RefreshToken,
		ExpiresIn:    authToken.ExpiresIn,
	}, nil
}

// signOut implements Service.
func (s *service) signOut(ctx context.Context, userID string, refreshTokenPayload *refreshTokenRequest) error {
	err := s.authManager.SignOut(ctx, userID, refreshTokenPayload.RefreshToken)

	if err != nil {
		return errors.Wrap(err, "s.authManager.SignOut")
	}

	return nil
}

// signUp implements Service.
func (s *service) signUp(ctx context.Context, payload *signUpRequest) (*signUpResponse, error) {
	userID, err := s.authManager.SignUp(ctx, &pkg.SignUpRequest{
		FirstName:  payload.FirstName,
		LastName:   payload.LastName,
		Email:      payload.Email,
		Password:   payload.Password,
		Phone:      payload.Phone,
		BirthDate:  payload.BirthDate,
		Role:       payload.Role,
		NotifToken: payload.NotifToken,
	})

	if err != nil {
		return nil, errors.Wrap(err, "s.authManager.SignUp")
	}

	return &signUpResponse{UserID: userID}, nil
}

// updateAvatar implements Service.
func (s *service) updateAvatar(ctx context.Context, user *user, file pkg.File, fileInfo pkg.FileInfo) error {
	// save old photo url which is used for rollback
	oldPhotoURL := user.PhotoURL

	// decode image to get information ()
	img, format, err := image.Decode(io.LimitReader(file, _MAX_AVATAR_BYTES))

	if err != nil {
		return errors.Wrap(err, "image.Decode")
	}

	// create new file with size 400x400, algo catmullrom to make image beautiful
	img = imaging.Fill(img, 400, 400, imaging.Center, imaging.CatmullRom)

	// using go nanoid to generate new name for file (for avatar)
	fileName, err := gonanoid.New(21)

	if err != nil {
		return errors.Wrap(err, "gonanoid.New")
	}

	buffer := new(bytes.Buffer)

	switch format {
	case "png":
		fileName += ".png"
		err = png.Encode(buffer, img)
	case "jpeg":
		fileName += ".jpeg"
		err = jpeg.Encode(buffer, img, nil)
	default:
		return errors.New("UnsupportedAvatarFormat: " + format)
	}

	// check error after encode
	if err != nil {
		return errors.Wrap(err, "format.Encode")
	}

	// get byte[] (byte of image)
	data := buffer.Bytes()

	// Upload image to S3 storage
	photoURL, err := s.storage.Upload(ctx, pkg.UploadInput{
		File:        bytes.NewReader(data),
		Name:        fileName,
		Size:        int64(len(data)),
		ContentType: fileInfo.ContentType(),
	})

	// handle err when upload to s3 fail
	if err != nil {
		return errors.Wrap(err, "s.storage.Upload")
	}

	// update avatar user
	err = s.authManager.UpdateUser(ctx, user.AccessToken, &pkg.UpdateUserRequest{
		ID:       user.ID,
		PhotoURL: &photoURL,
	})

	if err != nil {
		err = errors.Wrap(err, "s.authManager.UpdateUser")

		// if update user fail -> remove object from s3
		err2 := s.storage.Delete(ctx, photoURL)

		if err2 != nil {
			err2 = errors.Wrap(err2, "s.storage.Delete")
			err = errors.Wrap(err, err2.Error())
		}

		return err
	}

	// delete old photo (old object on s3)
	if oldPhotoURL != "" {
		err = s.storage.Delete(ctx, oldPhotoURL)

		if err != nil {
			return errors.Wrap(err, "s.storage.Delete")
		}
	}

	return nil
}

// updateUser implements Service.
func (s *service) updateUser(ctx context.Context, user *user, payload *updateUserRequest) error {
	userToUpdate := &pkg.UpdateUserRequest{
		ID: user.ID,
	}

	if payload.FirstName != "" {
		userToUpdate.FirstName = &payload.FirstName
	}

	if payload.LastName != "" {
		userToUpdate.LastName = &payload.LastName
	}

	if payload.Email != "" {
		userToUpdate.Email = &payload.Email
	}

	if payload.Phone != "" {
		userToUpdate.Phone = &payload.Phone
	}

	if payload.BirthDate != "" {
		userToUpdate.BirthDate = &payload.BirthDate
	}

	if payload.NotifToken != "" {
		userToUpdate.NotifToken = &payload.NotifToken
	}

	err := s.authManager.UpdateUser(ctx, user.AccessToken, userToUpdate)

	if err != nil {
		return errors.Wrap(err, "s.authManager.UpdateUser")
	}

	return nil
}
