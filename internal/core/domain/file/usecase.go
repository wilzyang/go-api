package file

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/wilzyang/go-api/pkg"
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

func (u UseCase) DoUpload(ctx context.Context, form multipart.File, filename string) (result Result, err error) {

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

	url := fmt.Sprintf("https://storage.cloud.google.com/%s/%s", att.Bucket, pkg.GenerateUrlEncode(filename))

	data := FileList{
		Filename:     filename,
		Size:         int64(att.Size),
		MediaLink:    url,
		DownloadLink: att.MediaLink,
		FileType:     att.ContentType,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err = u.fileRepo.DoInsertData(filename, data)

	if err != nil {
		e := fmt.Sprintf("[Domain] Insert Data : %v", err)
		return result, errors.New(e)
	}

	result.IsError = false
	result.Data = fmt.Sprintf("Success upload file %s", filename)

	return result, nil
}

func (u UseCase) DoDelete(ctx context.Context, fileid string) (result Result, err error) {
	d, err := u.fileRepo.DoGetFile("file_lists", map[string]interface{}{"id": fileid})
	if err != nil {
		e := fmt.Sprintf("[Domain] Get data : %v", err)
		return result, errors.New(e)
	}

	err = u.fileApi.DeleteFile(context.Background(), d.Filename)
	if err != nil {
		e := fmt.Sprintf("[Domain] Get File Google Cloud Store : %v", err)
		return result, errors.New(e)
	}

	result.IsError = false
	result.Data = fmt.Sprintf("Success delete file %s", d.Filename)

	return result, nil
}
