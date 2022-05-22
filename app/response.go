package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

const (
	Success = "success"
	Fail    = "fail"
	Err     = "error"
)

type Response struct {
	IsError bool        `json:"is_error"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func RespondError(g *gin.Context, err error) error {
	if webErr, ok := errors.Cause(err).(Error); ok {
		if webErr.Code == Internal {
			g.JSON(http.StatusInternalServerError, Response{
				IsError: true,
				Message: ErrInternal.Error(),
				Code:    http.StatusInternalServerError,
			})
			return err
		}
		if webErr.Code == Notfound {
			g.JSON(http.StatusNotFound, Response{
				IsError: true,
				Message: webErr.Message,
				Code:    http.StatusNotFound,
			})
			return err
		}
		if webErr.Code == InvalidData {
			g.JSON(http.StatusNotFound, Response{
				IsError: true,
				Message: webErr.Message,
				Code:    http.StatusBadRequest,
			})
			return err
		}
		//default error message if catch
		g.JSON(http.StatusBadRequest, Response{
			IsError: true,
			Message: webErr.Message,
			Code:    http.StatusBadRequest,
		})
		return err
	}
	//default other errors
	g.JSON(http.StatusInternalServerError, Response{
		IsError: true,
		Message: ErrInternal.Error(),
		Code:    http.StatusInternalServerError,
	})
	return err

}

func RespondSuccess(g *gin.Context, res Response) error {
	//default other errors
	g.JSON(http.StatusOK, Response{
		IsError: false,
		Message: res.Message,
		Code:    http.StatusOK,
		Data:    res.Data,
	})
	return nil
}
