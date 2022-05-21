package box

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"
)

type UseCase struct {
	boxRepo boxRepository
	boxApi  boxApi
	boxJWT  boxJWT
}

func NewBoxUseCase(boxRepo boxRepository, boxApi boxApi, boxJWT boxJWT) UseCase {
	return UseCase{
		boxRepo: boxRepo,
		boxApi:  boxApi,
		boxJWT:  boxJWT,
	}
}

func (u UseCase) Upload(ctx context.Context, form multipart.File, filename string) (result Result, err error) {

	var b BoxResponse
	var s Entries
	token, err := u.boxJWT.GenerateBoxJWT()
	result.IsError = true
	if err != nil {
		e := fmt.Sprintf("[Domain] Generate Token : %v", err)
		return result, errors.New(e)
	}

	ur, err := u.boxApi.UploadBox(context.Background(), form, filename, token)
	if err != nil {
		e := fmt.Sprintf("[Domain] Upload box : %v", err)
		return result, errors.New(e)
	}

	err = json.Unmarshal(ur, &b)

	if err != nil {
		e := fmt.Sprintf("[Domain] Unmarshal ur : %v", err)
		return result, errors.New(e)
	}

	gl, err := u.boxApi.GetSharedLink(context.Background(), b.Entries[0].Id, token)

	if err != nil {
		e := fmt.Sprintf("[Domain] Get shared link : %v", err)
		return result, errors.New(e)
	}

	err = json.Unmarshal(gl, &s)

	if err != nil {
		e := fmt.Sprintf("[Domain] Unmarshal gl : %v", err)
		return result, errors.New(e)
	}

	err = u.boxRepo.DoInsertData(s.Name, s.Size, s.SharedLink.Url, s.Id)

	if err != nil {
		e := fmt.Sprintf("[Domain] Insert Data : %v", err)
		return result, errors.New(e)
	}

	result.IsError = false

	return result, nil
}

func (u UseCase) Delete(ctx context.Context, filekey string) (result Result, err error) {

	//get filekey from delete params
	// fk := r.URL.Query()["fileKey"]
	// filekey := string(fk[0])
	result.IsError = true
	token, err := u.boxJWT.GenerateBoxJWT()
	result.IsError = true
	if err != nil {
		e := fmt.Sprintf("[Domain] Generate Token : %v", err)
		return result, errors.New(e)
	}
	//check filekey from database ->get box_file_id
	data, err := u.boxRepo.DoCheckIdByFilekey(filekey)
	if err != nil {
		e := fmt.Sprintf("[Domain] Check Filekey : %v", err)
		return result, errors.New(e)
	}

	//delete to box.com using box_file_id
	err = u.boxApi.DeleteBox(context.Background(), data.FileId, token)

	if err != nil {
		e := fmt.Sprintf("[Domain] Box Delete : %v", err)
		return result, errors.New(e)
	}

	err = u.boxRepo.DoDeleteByFilekey(filekey)
	//delete to database using filekey if data successfully deleted
	if err != nil {
		e := fmt.Sprintf("[Domain] Database Delete : %v", err)
		return result, errors.New(e)
	}

	result.IsError = false

	return result, err
}
