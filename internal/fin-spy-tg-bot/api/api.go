package api

import (
	"net/http"

	"github.com/Kodik77rus/fin-spy-tg-bot/storage"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

//Base API server instance
type API struct {
	//unexported fields
	config   *Config
	logger   *logrus.Logger
	tgbotapi *tgbotapi.BotAPI

	storage *storage.Storage
}

//API constructor
func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
	}
}

func (api *API) Start() error {
	defer closeDb(api.storage)

	//trying configurate logruss
	if err := api.setLogLevel(); err != nil {
		api.logger.Warnf("Logrus failed: %v \n Set default loglvl \"Debug\"", err)
	}

	//trying configurate store
	api.configureStore()

	if err := api.startTgBotApi(); err != nil {
		api.logger.Errorf("Telegram API server: %v", err)
	}

	api.logger.Info("server starting at port", api.config.Port)
	return http.ListenAndServe(api.config.Port, nil)
}

//set api.storage and connect to db
func (a *API) configureStore() error {
	st := storage.New(a.config.DatabaseURL, a.config.Storage)

	if err := st.Open(); err != nil {
		return err
	}

	a.storage = st
	return nil
}

//set api.logger as logrus log lvl
func (a *API) setLogLevel() error {
	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err //
	}
	a.logger.SetLevel(log_level)
	return nil
}

//set api.tgBot and connect to the telegram API server
func (a *API) startTgBotApi() error {
	bot, err := tgbotapi.NewBotAPI(a.config.TgBot)
	if err != nil {
		return err
	}
	a.tgbotapi = bot
	return nil
}

func closeDb(s *storage.Storage) {
	s.Close()
}
