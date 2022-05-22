package app

import "github.com/wilzyang/go-api/domain/file"

type AppModule struct {
	FileModule FileModule
}

func NewAppModule(
	fileModule FileModule,
) AppModule {
	return AppModule{
		FileModule: fileModule,
	}
}

type FileModule struct {
	FileIP file.InputPort
}
