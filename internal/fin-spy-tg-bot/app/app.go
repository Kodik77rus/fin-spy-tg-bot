package app

import (
	"github.com/Kodik77rus/fin-spy-tg-bot/internal/fin-spy-tg-bot/cron"
	"github.com/Kodik77rus/fin-spy-tg-bot/internal/fin-spy-tg-bot/telegram"
	"github.com/Kodik77rus/fin-spy-tg-bot/storage"
	"github.com/sirupsen/logrus"
)

//Base app server instance
type APP struct {
	config  *Config          //App config
	logger  *logrus.Logger   //logger
	bot     *telegram.Bot    //telegram bot
	storage *storage.Storage //db
	cron    *cron.Cron       //scheduler
}

//App constructor
func New(config *Config) *APP {
	return &APP{
		config: config,
		logger: logrus.New(),
	}
}

//Start App
func (app *APP) Start() error {
	defer closeDbConnection(app)

	app.setLogLevel()
	app.setStorage()
	app.setTgBotApp()

	app.logger.Info("App working!")

	app.bot.Start()
	return nil
}
