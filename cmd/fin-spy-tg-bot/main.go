package main

import (
	"github.com/Kodik77rus/fin-spy-tg-bot/internal/fin-spy-tg-bot/app"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"os"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		logrus.Warn("No .env file found, trying to create file")

		env := [3]string{"TG_BOT", "LOG_LVL", "DB_URL"}

		for _, i := range env {
			if _, exists := os.LookupEnv(i); !exists {
				logrus.Fatal("No .env var")
			}
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
