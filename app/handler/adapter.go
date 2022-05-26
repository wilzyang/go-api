package handler

import (
	"context"
	"mime/multipart"

	"github.com/wilzyang/go-api/internal/core/domain/file"
)

type Adapter struct {
	fileIP file.InputPort
}

func NewAdapter(fileIP file.InputPort, adt ...interface{}) Adapter {
	return Adapter{
		fileIP: fileIP,
	}
}

func (a Adapter) doUpload(ctx context.Context, form multipart.File, filename string) (result file.Result, err error) {
	return a.fileIP.DoUpload(ctx, form, filename)
}

func (a Adapter) doDelete(ctx context.Context, filename string) (result file.Result, err error) {
	return a.fileIP.DoDelete(ctx, filename)
}
