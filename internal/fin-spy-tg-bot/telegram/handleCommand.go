package telegram

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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
	isUser, _ := b.storage.FindUser(uint(message.From.ID))
	if isUser.RowsAffected == 1 { //if find user
		msg := massegaConstructor(message, fmt.Sprintf("Hello %s!", message.From.FirstName))
		if _, err := b.bot.Send(msg); err != nil {
			panic(err)
		}
		return nil
	}

	switch message.From.LanguageCode {
	case ru:
		return b.createUser(message, b.config.RuDictionary)
	case en:
		return b.createUser(message, b.config.EnDictionary)
	default:
		return b.chooseLanguage(message, b.config.EnDictionary)
	}
}

func (b *Bot) marketCommand(message *tgbotapi.Message, flags []string) error {
	switch flags[0] {
	case "show":
		if len(flags) == 3 && flags[1] == "all" {
			switch flags[2] {
			case "locations":
				return b.findMarketsWithParam(message, "location")
			case "countries":
				return b.findMarketsWithParam(message, "country")
			case "cities":
				return b.findMarketsWithParam(message, "city")
			default:
				return b.unknownMessage(message)
			}
		}

		switch flags[1] {
		case "all":
			markets, _ := b.storage.GetAllMarkets(1) //firts page

			for _, m := range markets.Markets {
				msg := massegaConstructor(message, *textParser(m))
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

func (b *Bot) createUser(message *tgbotapi.Message, dictionary interface{}) error {
	user.Id = uint(message.From.ID)
	user.UserName = message.From.FirstName
	user.Language = message.From.LanguageCode

	switch d := dictionary.(type) {
	case RuDictionary:
		//Find user in db and update user language
		if err := b.storage.CreateUser(&user); err != nil { //
			return err
		}

		msg := massegaConstructor(message, fmt.Sprintf("%s, %s", message.From.FirstName, d.startMessage))
		if _, err := b.bot.Send(msg); err != nil {
			panic(err)
		}
		return nil
	case EnDictionary:
		if err := b.storage.CreateUser(&user); err != nil { //
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

func (b *Bot) chooseLanguage(message *tgbotapi.Message, dictionary interface{}) error {
	msg := massegaConstructor(message, "Choose language")
	msg.ReplyMarkup = inlineKeyBoardConstructor("", "") //crutch

	if _, err := b.bot.Send(msg); err != nil {
		panic(err)
	}
	return nil
}

func (b *Bot) findMarketsWithParam(message *tgbotapi.Message, param string) error {
	res, _ := b.storage.FindMarketsWithParam(param)
	switch param {
	case "location":
		var location location

		location.location = res

		msg := massegaConstructor(message, *textParser(location))
		if _, err := b.bot.Send(msg); err != nil {
			panic(err)
		}
		return nil
	case "country":
		var country country

		country.country = res

		msg := massegaConstructor(message, *textParser(country))
		if _, err := b.bot.Send(msg); err != nil {
			panic(err)
		}
		return nil

	// 	return nil
	case "city":
		var city city

		city.city = res

		msg := massegaConstructor(message, *textParser(city))
		if _, err := b.bot.Send(msg); err != nil {
			panic(err)
		}
		return nil
	}
	return nil
}
