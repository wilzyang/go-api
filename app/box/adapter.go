package box

import (
	"context"
	"mime/multipart"

	"github.com/wilzyang/go-api/domain/box"
)

type Adapter struct {
	boxIP box.InputPort
}

func NewAdapter(boxIP box.InputPort, adt ...interface{}) Adapter {
	return Adapter{
		boxIP: boxIP,
	}
}

func (a Adapter) doUpload(ctx context.Context, form multipart.File, filename string) (result box.Result, err error) {
	return a.boxIP.Upload(ctx, form, filename)
}

func (a Adapter) doDelete(ctx context.Context, filename string) (result box.Result, err error) {
	return a.boxIP.Delete(ctx, filename)
}
