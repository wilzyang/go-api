package main

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/wilzyang/go-api/config/config"
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
	err := viper.Unmarshal(&conf)
	if err != nil {
		errors.Wrap(err, "Fail to unmarshal connections")
	}

	fmt.Println(conf)

	return err
}
