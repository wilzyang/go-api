package cli

import (
	"cloud.google.com/go/storage"
	"github.com/wilzyang/go-api/app"
	"github.com/wilzyang/go-api/domain/file"
	box "github.com/wilzyang/go-api/services/box"
	gcp "github.com/wilzyang/go-api/services/gcp"
	"gorm.io/gorm"

	repo "github.com/wilzyang/go-api/repo/file"
)

type BootstrapConfig struct {
	DB     *gorm.DB
	Debug  bool
	Client *storage.Client
	Bucket string
	BoxAPI Url
	BoxJWT box.BoxConfig
}

type Url struct {
	GenURL    string
	UploadURL string
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
