package box

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
)

type parent struct {
	Id string `json:"id"`
}

type attributes struct {
	Name   string `json:"name"`
	Parent parent `json:"parent"`
}

type UploadRes struct {
	Filename string `json:"file_name"`
	Data     []byte `json:"data"`
}

type ApiUrl struct {
	Upload string
	File   string
}

func NewBoxApi(a ApiUrl) *ApiUrl {
	return &ApiUrl{
		Upload: a.Upload,
		File:   a.File,
	}
}

func (u ApiUrl) UploadBox(f multipart.File, filename string, token string) ([]byte, error) {
	url := u.Upload
	method := "POST"
	//payload
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	attributes := &attributes{
		Name: filename,
		Parent: parent{
			Id: "0", //root folder is 0
		},
	}
	a, err := json.Marshal(attributes)
	if err != nil {
		e := fmt.Sprintf("Error marshaling attributes : %v", err)
		return nil, errors.New(e)
	}

	_ = writer.WriteField("attributes", string(a))

	//fmt.Println(string(a))

	uploadFile, _ := writer.CreateFormFile("file", filepath.Base(filename))

	// Copy the file to the destination path
	_, err = io.Copy(uploadFile, f)
	if err != nil {
		e := fmt.Sprintf("Error copy file : %v", err)
		return nil, errors.New(e)
	}

	err = writer.Close()
	if err != nil {
		e := fmt.Sprintf("Error writer close : %v", err)
		return nil, errors.New(e)
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		e := fmt.Sprintf("Error create request : %v", err)
		return nil, errors.New(e)
	}
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		e := fmt.Sprintf("Error send request : %v", err)
		return nil, errors.New(e)
	}

	//handling if get any error response, will return error
	if res.StatusCode < 200 || res.StatusCode > 299 {
		e := fmt.Sprintf("Got Error Response: %d", res.StatusCode)
		return nil, errors.New(e)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		e := fmt.Sprintf("Error read body request : %v", err)
		return nil, errors.New(e)
	}

	return body, err
}

func (u ApiUrl) GetSharedLink(token string, fileid string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", u.File, fileid)
	method := "PUT"

	payload := `{
		"shared_link": {
		  "access": "open",
		  "permissions": {
			"can_preview": true,
	  		"can_download": true
		  }
		}
	  }`

	client := &http.Client{}
	req, err := http.NewRequest(method, url, strings.NewReader(payload))

	if err != nil {
		e := fmt.Sprintf("Error create request : %v", err)
		return nil, errors.New(e)
	}

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		e := fmt.Sprintf("Error send request : %v", err)
		return nil, errors.New(e)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		e := fmt.Sprintf("Error read body request : %v", err)
		return nil, errors.New(e)
	}

	return body, err

}

func (u ApiUrl) DeleteBox(token string, fileid string) error {
	url := fmt.Sprintf("%s/%s", u.File, fileid)
	method := "DELETE"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		e := fmt.Sprintf("Error create request : %v", err)
		return errors.New(e)
	}

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		e := fmt.Sprintf("Error send request : %v", err)
		return errors.New(e)
	}

	//check if response not success
	if res.StatusCode < 200 || res.StatusCode > 299 {
		e := fmt.Sprintf("Get error from box: %v", res.StatusCode)
		return errors.New(e)
	}

	defer res.Body.Close()

	return err
}
