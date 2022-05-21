package router

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/wilzyang/go-api/app"
	"github.com/wilzyang/go-api/app/box"
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

	boxRoutes := box.NewBoxRoute(appModule.BoxModule.BoxIP)
	ApiGroup.POST(box.MainPath, boxRoutes.DoUpload)
	ApiGroup.DELETE(box.MainPath, boxRoutes.DoDelete)

	return g, nil
}
