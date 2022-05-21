package box

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/wilzyang/go-api/app"
	"github.com/wilzyang/go-api/domain/box"
)

const (
	MainPath = "/files"
)

type Routes struct {
	boxIP box.InputPort
}

type Response struct {
	IsError bool        `json:"is_error"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewBoxRoute(boxIP box.InputPort) Routes {
	return Routes{
		boxIP: boxIP,
	}
}

func (r Routes) DoUpload(c *gin.Context) {

	adapter := NewAdapter(r.boxIP)

	err := c.Request.ParseMultipartForm(32 << 20)

	if err != nil {
		app.RespondError(c, err)
	}

	f, h, err := c.Request.FormFile("file")
	if err != nil {
		app.RespondError(c, err)
	}

	defer f.Close()

	_, err = adapter.doUpload(context.Background(), f, h.Filename)

	if err != nil {
		app.RespondError(c, app.Error{
			Code: app.Internal,
		})
	}

	app.RespondSuccess(c, app.Response{
		Message: "Upload Success",
	})

}

func (r Routes) DoDelete(c *gin.Context) {

	adapter := NewAdapter(r.boxIP)

	filekey := c.Query("fileKey")

	_, err := adapter.doDelete(context.Background(), filekey)

	if err != nil {
		app.RespondError(c, app.Error{
			Code: app.Internal,
		})
	}

	app.RespondSuccess(c, app.Response{
		Message: "Delete Success",
	})

}
