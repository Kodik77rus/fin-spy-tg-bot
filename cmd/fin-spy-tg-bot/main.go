package main

import (
	"flag"
	"os"
	"runtime"

	"github.com/BurntSushi/toml"
	"github.com/Kodik77rus/fin-spy-tg-bot/internal/fin-spy-tg-bot/app"
	"github.com/sirupsen/logrus"
)

var (
	configPath string
)

func init() {
	//route to the config file
	flag.StringVar(&configPath, "path", "../../configs/app.toml", "path to config file in .toml format")
}

func main() {
	runtime.GOMAXPROCS(4)

	flag.Parse()

	//create base config
	config := app.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		logrus.Warnf("Can't find config file! Using default values: %v", err)
	}

	//set config from toml file
	server := app.New(config)

	//app server start
	if err := server.Start(); err != nil {
		os.Exit(1)
	}

}
