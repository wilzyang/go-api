package config_sample

import (
	"flag"
	"io/ioutil"
	"log"
	"strings"

	"github.com/spf13/viper"
)

var (
	Debug bool
)

func ReadConfig() {
	debug := flag.Bool("debug", Debug, "debug mode")
	flag.Parse()
	Debug = *debug

	files, _ := ioutil.ReadDir("config/")

	viper.AddConfigPath("config/")
	for _, file := range files {
		name := strings.Split(file.Name(), ".")[0]
		viper.SetConfigName(name)
		//handle all config
		err := viper.MergeInConfig()
		if err != nil {
			log.Fatalf("[Server] reading config file %s \n", err)
		}
	}
}
