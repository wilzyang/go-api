package file

import (
	"context"
	"mime/multipart"
)

type InputPort interface {
	DoUpload(ctx context.Context, form multipart.File, filename string) (result Result, err error)
	DoDelete(ctx context.Context, fileid string) (result Result, err error)
}
