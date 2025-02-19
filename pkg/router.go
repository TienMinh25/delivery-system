package pkg

import "io"

type File interface {
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Closer
}

type FileInfo interface {
	FileName() string
	FileSize() int64
	ContentType() string
}