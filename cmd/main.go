package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"cloud.google.com/go/storage"
	"github.com/apex/log"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/wilzyang/go-api/app/router"
	"github.com/wilzyang/go-api/cmd/cli"
	"github.com/wilzyang/go-api/config"
	repo "github.com/wilzyang/go-api/internal/repository"
)

var (
	conf Config
)

func init() {
	config.ReadConfig()
}

func main() {
	if err := run(); err != nil {
		log.Fatalf("[Server] %s \n", err)
	}
}

func run() error {

	//get main config
	err := viper.Unmarshal(&conf)
	if err != nil {
		errors.Wrap(err, "Fail to unmarshal connections")
	}

	//database connection
	conn, err := repo.ConnectPsql(conf.Database.Dsn)

	if err != nil {
		return errors.Wrap(err, "Connect database")
	}

	//GCP client connection
	// Creates a client.
	client, err := storage.NewClient(context.Background())
	if err != nil {
		return errors.Wrap(err, "Failed to create client")
	}
	defer client.Close()

	bc := cli.BootstrapConfig{
		DB:     conn,
		Client: client,
		Bucket: conf.Gcp.Bucket,
		BoxAPI: cli.BoxApi{
			Files:  conf.BoxApi.Files,
			Upload: conf.BoxApi.Upload,
		},
	}

	appModule, err := cli.Bootstrap(bc)

	if err != nil {
		return errors.Wrap(err, "Bootstraping")
	}

	c, err := router.Routes(appModule)

	if err != nil {
		return errors.Wrap(err, "Creating Routes")
	}

	//Server
	srv := &http.Server{
		Addr:    conf.Server.Address,
		Handler: c,
	}

	go func() {
		log.Warnf("[Server] Starting the apps on port %s \n", conf.Server.Address)
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("[Server] Shutting the apps... [%s]", err)
		}
	}()

	//Graceful Shutdown
	timeout := 1 * time.Minute
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}
