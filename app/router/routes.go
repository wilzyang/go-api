package router

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/wilzyang/go-api/app"
	"github.com/wilzyang/go-api/app/handler"
)

func Routes(appModule app.AppModule) (*gin.Engine, error) {
	g := NewAPI()

	g.Use(
		gin.Recovery(),
		gzip.Gzip(gzip.DefaultCompression),
		CORSMiddleware(),
	)
	gin.SetMode(gin.ReleaseMode)

	ApiGroup := g.Group("/api")

	fileRoutes := handler.NewFileRoute(appModule.FileModule.FileIP)
	ApiGroup.POST(handler.MainPath, fileRoutes.DoUpload)
	ApiGroup.DELETE(handler.MainPath, fileRoutes.DoDelete)
	g.NoRoute(handler.NotFound)

	return g, nil
}
