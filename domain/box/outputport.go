package box

import (
	"context"
	"mime/multipart"
)

type boxRepository interface {
	DoInsertData(title string, size int, link string, file_id string) (err error)
	DoCheckIdByFilekey(filekey string) (data ResearchReport, err error)
	DoDeleteByFilekey(filekey string) (err error)
}

type boxJWT interface {
	GenerateBoxJWT() (string, error)
}

type boxApi interface {
	UploadBox(ctx context.Context, f multipart.File, filename string, token string) ([]byte, error)
	GetSharedLink(ctx context.Context, fileid string, token string) ([]byte, error)
	DeleteBox(ctx context.Context, fileid string, token string) error
}
