package telegram

import (
	"fmt"
	"github.com/Kodik77rus/fin-spy-tg-bot/internal/fin-spy-tg-bot/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"strings"
)

const (
	//commands
	commandStart  = "start"
	commandMarket = "markets"
	// commandWhatch     = "whatch"
	// commandDelete     = "delete"
	// commandWhatchList = "whatchlist"
	// commandInfo       = "info"

	//collback message data
	ru   = "ru"
	en   = "en"
	page = "page"
)

var user models.User

//Handle commands
func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	command := strings.Split(message.Command(), "_")

	switch command[0] {
	case commandStart:
		return b.startCommand(message)
	case commandMarket:
		return b.marketCommand(message, command[1:])
	default:
		return b.unknownCommand(message)
	}
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

	data := strings.Split(cb.Data, "=")
	switch data[0] {
	case ru:
		//Find user in db and update user language
		if err := b.storage.UpdateUser(&user); err != nil {
			return err
		}

		msg := massegaConstructor(cb.Message, "RU")

		if _, err := b.bot.Send(msg); err != nil {
			panic(err)
		}

	case en:
		//Find user in db and update user language
		if err := b.storage.UpdateUser(&user); err != nil {
			return err
		}

		msg := massegaConstructor(cb.Message, "EN")

		if _, err := b.bot.Send(msg); err != nil {
			panic(err)
		}
	case page:
		params := strings.Split(cb.Data, ",")

		p := paginationParser(params)
		if !p.isValid {
			msg := massegaConstructor(cb.Message, "Bad query")
			if _, err := b.bot.Send(msg); err != nil {
				panic(err)
			}
		}

		switch p.query {
		case "all_markets":
			markets, _ := b.storage.GetAllMarkets(p.page + 1) //next page
			if markets.Count == 0 {
				msg := massegaConstructor(cb.Message, "You watched all markets!")
				if _, err := b.bot.Send(msg); err != nil {
					panic(err)
				}
				return nil
			}

			for _, m := range markets.Markets {
				parsedTxt := textParser(m)

				msg := massegaConstructor(cb.Message, parsedTxt)
				msg.ReplyMarkup = inlineKeyBoardConstructor("info", m.Hour)

				if _, err := b.bot.Send(msg); err != nil {
					panic(err)
				}
			}
			msg := massegaConstructor(cb.Message, "Touch to see next markets")
			msg.ReplyMarkup = inlineKeyBoardConstructor("next", fmt.Sprintf("page=%d,query=%s", p.page+1, p.query))
			if _, err := b.bot.Send(msg); err != nil {
				panic(err)
			}
			return nil
		}
	default:
		return b.unknownCommand(cb.Message)
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
		msg.ReplyMarkup = inlineKeyBoardConstructor("", "") //crutch

		if _, err := b.bot.Send(msg); err != nil {
			panic(err)
		}
	}
	return nil
}

func (b *Bot) marketCommand(message *tgbotapi.Message, flags []string) error {
	switch flags[0] {
	case "show":
		switch flags[1] {
		case "all":
			markets, _ := b.storage.GetAllMarkets(1) //firts page
			for _, m := range markets.Markets {
				parsedTxt := textParser(m)
				msg := massegaConstructor(message, parsedTxt)
				msg.ReplyMarkup = inlineKeyBoardConstructor("info", m.Hour)
				if _, err := b.bot.Send(msg); err != nil {
					panic(err)
				}
			}
			msg := massegaConstructor(message, "Touch to see next markets")
			msg.ReplyMarkup = inlineKeyBoardConstructor("next", "page=1,query=all_markets")
			if _, err := b.bot.Send(msg); err != nil {
				panic(err)
			}
			return nil
		case "code":

		case "mic":

		case "location":

		case "country":

		case "city":

		case "delay":
		default:
			return b.unknownCommand(message)
		}
	case "subscribe":
	default:
		return b.unknownCommand(message)
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
