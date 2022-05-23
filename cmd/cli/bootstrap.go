package cli

import (
	"cloud.google.com/go/storage"
	"github.com/wilzyang/go-api/app"
	"github.com/wilzyang/go-api/internal/core/domain/file"
	"github.com/wilzyang/go-api/internal/core/services/gcp"
	"gorm.io/gorm"

	repo "github.com/wilzyang/go-api/internal/repository/file"
)

type BootstrapConfig struct {
	DB     *gorm.DB
	Debug  bool
	Client *storage.Client
	Bucket string
	BoxAPI BoxApi
	BoxJWT string
}

type BoxApi struct {
	Files  string
	Upload string
}

func Bootstrap(config BootstrapConfig) (app.AppModule, error) {

	//pass repo configuration
	fileRepo := repo.NewFileRepository(config.DB)

	//pass gcp client configuration
	client := gcp.GcpClient{
		Client: config.Client,
		Bucket: config.Bucket,
	}

	gcp := gcp.NewGcpAPi(client)

	fileIP := file.NewFileUseCase(fileRepo, gcp)

	//pass fileIP that already fill in with configuration to File Module
	fileModule := app.FileModule{
		FileIP: fileIP,
	}

	appModule := app.AppModule{
		FileModule: fileModule,
	}

	return appModule, nil

}
