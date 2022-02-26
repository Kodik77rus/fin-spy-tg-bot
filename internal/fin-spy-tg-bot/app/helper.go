package app

import (
	"github.com/Kodik77rus/fin-spy-tg-bot/internal/fin-spy-tg-bot/cron"
	"github.com/Kodik77rus/fin-spy-tg-bot/internal/fin-spy-tg-bot/telegram"
	"github.com/Kodik77rus/fin-spy-tg-bot/storage"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

//Set app.logger as logrus log lvl
func (app *APP) setLogLevel() error {
	log_level, err := logrus.ParseLevel(app.config.LoggerLevel)
	if err != nil {
		app.logger.Warn("Using default log lvl: \"debug\"")
		return err
	}

	app.logger.SetFormatter(&logrus.JSONFormatter{})
	app.logger.SetLevel(log_level)

	return nil
}

//Set app.storage and connect to db
func (app *APP) setStorage() error {
	st := storage.New(app.config.DatabaseURL, app.config.Storage)
	if err := st.Open(); err != nil {
		app.logger.Panicf("Connection to the database failed", err)
	}

	app.storage = st
	return nil
}

//Set app.bot and connect to the telegram server
func (app *APP) setTgBotApp() error {
	bot, err := tgbotapi.NewBotAPI(app.config.TgBot)
	if err != nil {
		app.logger.Errorf("Connection to the Telegram API server failed: %v", err)
	}

	bot.Debug = true

	crn := cron.New(app.config.FinhubToken)

	app.bot = telegram.New(bot, app.storage, crn)

	return nil
}

func closeDbConnection(app *APP) error {
	if err := app.storage.Close(); err != nil {
		app.logger.Panicf("Database connection is not closed: %v", err)
	}
	return nil
}
