package app

import "github.com/wilzyang/go-api/domain/box"

type AppModule struct {
	BoxModule BoxModule
}

func NewAppModule(
	boxModule BoxModule,
) AppModule {
	return AppModule{
		BoxModule: boxModule,
	}
}

type BoxModule struct {
	BoxIP box.InputPort
}
