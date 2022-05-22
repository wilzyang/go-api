package router

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/wilzyang/go-api/app"
	"github.com/wilzyang/go-api/app/file"
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

	fileRoutes := file.NewFileRoute(appModule.FileModule.FileIP)
	ApiGroup.POST(file.MainPath, fileRoutes.DoUpload)
	// ApiGroup.DELETE(file.MainPath, boxRoutes.DoDelete)

	return g, nil
}
