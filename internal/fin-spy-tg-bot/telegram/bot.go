package telegram

import (
	"github.com/Kodik77rus/fin-spy-tg-bot/internal/fin-spy-tg-bot/cron"
	"github.com/Kodik77rus/fin-spy-tg-bot/storage"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	config  *Config
	bot     *tgbotapi.BotAPI
	cron    *cron.Cron
	storage *storage.Storage
}

func New(bot *tgbotapi.BotAPI, storage *storage.Storage, cron *cron.Cron) *Bot {
	return &Bot{
		config:  NewConfig(),
		bot:     bot,
		storage: storage,
		cron:    cron,
	}
}

func (b *Bot) Start() error {
	//sceduler
	go b.cron.Start()

	// update listener
	updates := b.initUpdateChanel()

	b.handleUpdates(updates)

	return nil
}

func (b *Bot) initUpdateChanel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	return b.bot.GetUpdatesChan(u)
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.CallbackQuery != nil {
			b.callbackQueryHandler(update.CallbackQuery)
			continue
		}
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			b.handleCommand(update.Message)
		}
	}
}
