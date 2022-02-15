package telegram

import (
	"fmt"

	"github.com/Kodik77rus/fin-spy-tg-bot/internal/fin-spy-tg-bot/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"strings"
)

const (
	//collback message data
	ru   = "ru"
	en   = "en"
	page = "page"
)

var user models.User

//Handle callback querys
func (b *Bot) callbackQueryHandler(cb *tgbotapi.CallbackQuery) error {
	// Respond to the callback query, telling Telegram to show the user
	// a message with the data received.
	callback := tgbotapi.NewCallback(cb.ID, cb.Data)
	if _, err := b.bot.Request(callback); err != nil {
		panic(err)
	}

	data := strings.Split(cb.Data, "=")
	switch data[0] {
	case ru:
		return b.setUserLanguage(cb, b.config.RuDictionary)
	case en:
		return b.setUserLanguage(cb, b.config.EnDictionary)
	case page:
		params := strings.Split(cb.Data, ",")
		p := paginationParser(params)
		if !p.isValid {
			msg := massegaConstructor(cb.Message, "Bad query")
			return b.sendMessage(msg)
		}

		switch p.query {
		case "all_markets":
		case "market":
			return b.sendAllMarkets(cb.Message, p.page) //next page
		case "location", "country", "city":
			return b.FindMarketsWithParams(cb.Message, p)
		default:
			return b.unknownMessage(cb.Message)
		}

	default:
		return b.unknownMessage(cb.Message)
	}
	return nil
}

//Find user in db and update user language
func (b *Bot) setUserLanguage(cb *tgbotapi.CallbackQuery, dictionary interface{}) error {
	user.Id = uint(cb.Message.Chat.ID)
	user.Language = cb.Data

	switch d := dictionary.(type) {
	case RuDictionary:
		if err := b.storage.UpdateUser(&user); err != nil {
			return err
		}
		msg := massegaConstructor(cb.Message, fmt.Sprintf("%s, %s", cb.Message.From.FirstName, d.setLanguage))
		return b.sendMessage(msg)
	case EnDictionary:
		if err := b.storage.UpdateUser(&user); err != nil {
			return err
		}
		msg := massegaConstructor(cb.Message, fmt.Sprintf("%s, %s", cb.Message.From.FirstName, d.setLanguage))
		return b.sendMessage(msg)
	}
	return nil
}
