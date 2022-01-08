package app

import (
	"sync"

	"github.com/Kodik77rus/fin-spy-tg-bot/internal/fin-spy-tg-bot/telegram"
	"github.com/Kodik77rus/fin-spy-tg-bot/storage"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

//Prepair func for goroutine
type AppStarter struct {
	sas func(a *APP, wg *sync.WaitGroup) error
}

//App Starter constructor
func newAppStarter(f func(a *APP, wg *sync.WaitGroup) error) *AppStarter {
	return &AppStarter{sas: f}
}

//Call method for App Starter
func (g *AppStarter) start(a *APP, wg *sync.WaitGroup) error { return g.sas(a, wg) }

//Set app.logger as logrus log lvl
func setLogLevel(app *APP, wg *sync.WaitGroup) error {
	defer wg.Done()

	log_level, err := logrus.ParseLevel(app.config.LoggerLevel)
	if err != nil {
		app.logger.Warn("Using default log lvl: \"debug\"")
		return err
	}

	app.logger.SetLevel(log_level)
	return nil
}

//Set app.storage and connect to db
func setStorage(app *APP, wg *sync.WaitGroup) error {
	defer wg.Done()

	st := storage.New(app.config.DatabaseURL, app.config.Storage)
	if err := st.Open(); err != nil {
		app.logger.Panicf("Connection to the database failed", err)
	}

	app.storage = st
	return nil
}

//Set app.bot and connect to the telegram server
func setTgBotApp(app *APP, wg *sync.WaitGroup) error {
	defer wg.Done()
	bot, err := tgbotapi.NewBotAPI(app.config.TgBot)
	if err != nil {
		app.logger.Panicf("Connection to the Telegram API server failed: %v", err)
	}

	bot.Debug = true

	app.bot = telegram.New(bot)

	return nil
}

func closeDbConnection(app *APP) error {
	if err := app.storage.Close(); err != nil {
		app.logger.Panicf("Database connection is not closed: %v", err)
	}
	return nil
}
