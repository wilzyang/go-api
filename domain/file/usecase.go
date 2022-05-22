package file

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
)

type UseCase struct {
	fileRepo fileRepository
	fileApi  fileApi
}

func NewFileUseCase(fileRepo fileRepository, fileApi fileApi) UseCase {
	return UseCase{
		fileRepo: fileRepo,
		fileApi:  fileApi,
	}
}

func (u UseCase) Upload(ctx context.Context, form multipart.File, filename string) (result Result, err error) {

	err = u.fileApi.UploadFile(context.Background(), form, filename)
	if err != nil {
		e := fmt.Sprintf("[Domain] Upload Google Cloud Store : %v", err)
		return result, errors.New(e)
	}

	att, err := u.fileApi.GetFile(context.Background(), filename)
	if err != nil {
		e := fmt.Sprintf("[Domain] Get File Google Cloud Store : %v", err)
		return result, errors.New(e)
	}

	err = u.fileRepo.DoInsertData(filename, int(att.Size), att.MediaLink, att.ContentType)

	if err != nil {
		e := fmt.Sprintf("[Domain] Insert Data : %v", err)
		return result, errors.New(e)
	}

	result.IsError = false

	return result, nil
}
