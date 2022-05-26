package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wilzyang/go-api/app"
	"github.com/wilzyang/go-api/internal/core/domain/file"
)

const (
	MainPath = "/files"
	FilePath = "/files/:fileid"
)

type Routes struct {
	fileIP file.InputPort
}

type Response struct {
	IsError bool        `json:"is_error"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewFileRoute(fileIP file.InputPort) Routes {
	return Routes{
		fileIP: fileIP,
	}
}

func (r Routes) DoUpload(c *gin.Context) {

	adapter := NewAdapter(r.fileIP)

	err := c.Request.ParseMultipartForm(32 << 20)

	if err != nil {
		app.RespondError(c, err)
	}

	f, h, err := c.Request.FormFile("file")
	if err != nil {
		app.RespondError(c, err)
	}

	defer f.Close()

	data, err := adapter.doUpload(context.Background(), f, h.Filename)

	if err != nil {
		app.RespondError(c, err)
	} else {
		app.RespondSuccess(c, app.Response{
			IsError: false,
			Code:    http.StatusOK,
			Message: "Upload Success",
			Data:    data,
		})
	}
}

func (r Routes) DoDelete(c *gin.Context) {

	adapter := NewAdapter(r.fileIP)

	id := c.Param("fileid")

	data, err := adapter.doDelete(context.Background(), id)

	if err != nil {
		app.RespondError(c, err)
	} else {
		app.RespondSuccess(c, app.Response{
			IsError: false,
			Code:    http.StatusOK,
			Message: "Delete Success",
			Data:    data,
		})
	}

}
