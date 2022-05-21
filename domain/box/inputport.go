package box

import (
	"context"
	"mime/multipart"
)

type InputPort interface {
	Upload(ctx context.Context, form multipart.File, filename string) (result Result, err error)
	Delete(ctx context.Context, filename string) (result Result, err error)
}
