package main

import (
	"flag"

	"github.com/BurntSushi/toml"
	"github.com/Kodik77rus/fin-spy-tg-bot/internal/fin-spy-tg-bot/api"

	"log"
)

var (
	configPath string
)

func init() {
	//route file path to config file
	flag.StringVar(&configPath, "path", "../../configs/api.toml", "path to config file in .toml format")
}

func main() {
	flag.Parse()

	//create base config
	config := api.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Println("can not find configs file. using default values:", err)
	}

	//set config from toml file
	server := api.New(config)

	//api server start
	log.Fatal(server.Start())
}
