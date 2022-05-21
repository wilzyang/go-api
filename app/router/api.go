package router

import "github.com/gin-gonic/gin"

func NewAPI() *gin.Engine {
	g := gin.New()

	return g
}
