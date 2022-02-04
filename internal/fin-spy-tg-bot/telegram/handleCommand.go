package telegram

import (
	"fmt"
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

)

//Handle commands
func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	command := strings.Split(message.Command(), "_")

	switch command[0] {
	case commandStart:
		return b.startCommand(message)
	case commandMarket:
		return b.marketCommand(message, command[1:])
	default:
		return b.unknownMessage(message)
	}
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

	switch message.From.LanguageCode {
	case ru:
		return b.findUser(message, b.config.RuDictionary)
	case en:
		return b.findUser(message, b.config.EnDictionary)
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
			return b.unknownMessage(message)
		}
	case "subscribe":
	default:
		return b.unknownMessage(message)
	}
	return nil
}

func (b *Bot) findUser(message *tgbotapi.Message, dictionary interface{}) error {
	user.Id = uint(message.From.ID)
	user.UserName = message.From.FirstName
	user.Language = message.From.LanguageCode

	switch d := dictionary.(type) {
	case RuDictionary:
		//Find user in db and update user language
		if err := b.storage.CreateUser(&user); err != nil {
			return err
		}

		msg := massegaConstructor(message, fmt.Sprintf("%s, %s", message.From.FirstName, d.startMessage))
		if _, err := b.bot.Send(msg); err != nil {
			panic(err)
		}
		return nil
	case EnDictionary:
		if err := b.storage.CreateUser(&user); err != nil {
			return err
		}

		msg := massegaConstructor(message, fmt.Sprintf("%s, %s", message.From.FirstName, d.startMessage))
		if _, err := b.bot.Send(msg); err != nil {
			panic(err)
		}
		return nil
	}
	return nil
}
