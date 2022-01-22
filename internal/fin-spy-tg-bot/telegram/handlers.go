package telegram

import (
	"fmt"
	"strings"

	"github.com/Kodik77rus/fin-spy-tg-bot/internal/fin-spy-tg-bot/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	commandStart   = "start"
	commandMarkets = "markets"
	// commandWhatch     = "whatch"
	// commandDelete     = "delete"
	// commandWhatchList = "whatchlist"
	// commandInfo       = "info"

	ru = "ru"
	en = "en"
)

var user models.User

//Handle commands
func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	command := strings.Split(message.Command(), "_")

	switch command[0] {
	case commandStart:
		return b.startCommand(message)
	case commandMarkets:

	default:
		return b.unknownCommand(message)
	}
	return nil
}

//Handle callback querys
func (b *Bot) callbackQueryHandler(cb *tgbotapi.CallbackQuery) error {
	// Respond to the callback query, telling Telegram to show the user
	// a message with the data received.
	callback := tgbotapi.NewCallback(cb.ID, cb.Data)
	if _, err := b.bot.Request(callback); err != nil {
		panic(err)
	}

	user.Id = uint(cb.Message.Chat.ID)
	user.Language = cb.Data

	switch cb.Data {
	case ru:
		if err := b.setUserLanguage(&user); err != nil {
			return err
		}

		msg := massegaConstructor(cb.Message, "RU")
		if _, err := b.bot.Send(msg); err != nil {
			panic(err)
		}

	case en:
		if err := b.setUserLanguage(&user); err != nil {
			return err
		}

		msg := massegaConstructor(cb.Message, "EN")
		if _, err := b.bot.Send(msg); err != nil {
			panic(err)
		}
	}
	return nil
}

//Command start handler
func (b *Bot) startCommand(message *tgbotapi.Message) error {
	isUser, err := b.storage.FindUser(uint(message.From.ID))
	if err != nil {
		msg := massegaConstructor(message, "troubls with bd")
		if _, err := b.bot.Send(msg); err != nil {
			panic(err)
		}
		return err
	}

	//if find user
	if isUser.RowsAffected == 1 {
		msg := massegaConstructor(message, fmt.Sprintf("Hello %s!", message.From.FirstName))
		if _, err := b.bot.Send(msg); err != nil {
			panic(err)
		}
		return nil
	}

	user.Id = uint(message.From.ID)
	user.UserName = message.From.FirstName
	user.Language = message.From.LanguageCode

	switch message.From.LanguageCode {
	case ru:
		if err := b.storage.CreateUser(&user); err != nil {
			return err
		}

		msg := massegaConstructor(message, "RU")
		if _, err := b.bot.Send(msg); err != nil {
			panic(err)
		}

		return nil
	case en:
		if err := b.storage.CreateUser(&user); err != nil {
			return err
		}

		msg := massegaConstructor(message, "RU")
		if _, err := b.bot.Send(msg); err != nil {
			panic(err)
		}

		return nil
	default:
		msg := massegaConstructor(message, "Choose language")
		msg.ReplyMarkup = inlineLanguageKeyBoard

		if _, err := b.bot.Send(msg); err != nil {
			panic(err)
		}
	}
	return nil
}

//Send default message for unknown command
func (b *Bot) unknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Silly bot Finn don't understant you!")
	if _, err := b.bot.Send(msg); err != nil {
		panic(err)
	}
	return nil
}
