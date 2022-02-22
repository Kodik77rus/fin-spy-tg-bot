package main

import (
	"github.com/Kodik77rus/fin-spy-tg-bot/internal/fin-spy-tg-bot/app"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		logrus.Warn("No .env file found, trying to create file")

		file, err := os.Open(".env.example")
		if err != nil {
			logrus.Fatal("no config file")
		}

		env, err := godotenv.Parse(file)
		if err != nil {
			logrus.Fatal("can't parse file")
		}

		if err := godotenv.Write(env, ".env"); err != nil {
			logrus.Fatal("can't write file")
		}

		if err := godotenv.Load(); err != nil {
			logrus.Fatal(err)
		}
	}
}

func main() {
	//create base config
	config := app.NewConfig()

	//set config from toml file
	server := app.New(config)

	//app server start
	if err := server.Start(); err != nil {
		os.Exit(1)
	}

}
