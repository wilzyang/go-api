package handler

import (
	"github.com/gin-gonic/gin"
)

func NotFound(c *gin.Context) {
	c.JSON(404, Response{
		IsError: true,
		Code:    404,
		Message: "Page not found",
	})
}
