package file

import (
	"context"
	"mime/multipart"

	"cloud.google.com/go/storage"
)

type fileRepository interface {
	DoInsertData(title string, size int, link string, file_type string) (err error)
}

type fileApi interface {
	UploadFile(ctx context.Context, f multipart.File, filename string) error
	GetFile(ctx context.Context, object string) (*storage.ObjectAttrs, error)
}
