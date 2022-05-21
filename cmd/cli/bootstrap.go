package cli

import (
	"github.com/wilzyang/go-api/app"
	"github.com/wilzyang/go-api/domain/box"
	"gorm.io/gorm"

	repo "github.com/wilzyang/go-api/repo/box"
	service "github.com/wilzyang/go-api/services/box"
)

type BootstrapConfig struct {
	DB     *gorm.DB
	Debug  bool
	BoxAPI Url
	BoxJWT service.BoxConfig
}

type Url struct {
	GenURL    string
	UploadURL string
}

func Bootstrap(config BootstrapConfig) (app.AppModule, error) {
	boxRepo := repo.NewBoxRepository(config.DB)

	bc := service.BoxConfig{
		PublicKeyID:  config.BoxJWT.PublicKeyID,
		ClientID:     config.BoxJWT.ClientID,
		Sub:          config.BoxJWT.Sub,
		ClientSecret: config.BoxJWT.ClientSecret,
	}

	cred := service.NewBoxAuth(bc)

	boxIP := box.NewBoxUseCase(boxRepo, nil, cred)

	boxModule := app.BoxModule{
		BoxIP: boxIP,
	}

	appModule := app.AppModule{
		BoxModule: boxModule,
	}

	return appModule, nil

}
