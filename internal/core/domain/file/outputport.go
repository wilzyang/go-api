package file

import (
	"context"
	"mime/multipart"

	"cloud.google.com/go/storage"
)

type fileRepository interface {
	DoInsertData(title string, data FileList) (err error)
	DoGetFile(table string, condition map[string]interface{}) (data FileList, err error)
}

type fileApi interface {
	UploadFile(ctx context.Context, f multipart.File, filename string) error
	GetFile(ctx context.Context, object string) (*storage.ObjectAttrs, error)
	DeleteFile(ctx context.Context, object string) error
}
