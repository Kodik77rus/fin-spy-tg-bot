package telegram

import (
	"github.com/Kodik77rus/fin-spy-tg-bot/storage"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	config  *Config
	bot     *tgbotapi.BotAPI
	storage *storage.Storage
}

func New(bot *tgbotapi.BotAPI, storage *storage.Storage) *Bot {
	return &Bot{
		config:  NewConfig(),
		bot:     bot,
		storage: storage,
	}
}

func (b *Bot) Start() error {

	return nil
}
