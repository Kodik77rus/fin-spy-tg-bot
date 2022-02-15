package telegram

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	//commands

	commandStart  = "start"
	commandMarket = "market"

	// commandWhatch     = "whatch"
	// commandDelete     = "delete"
	// commandWhatchList = "whatchlist"
	// commandInfo       = "info"

)

var p pagination

//Handle commands
func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	command := commandValidation(strings.Split(message.Command(), "_"))

	switch command.name {
	case commandStart:
		return b.startCommand(message)
	case commandMarket:
		switch command.flag {
		case "show":
			if len(command.param) == 0 {
				return b.unknownMessage(message)
			}

			switch command.param[0] {
			case "all":
				return b.sendAllMarkets(message, 1) //firts page
			case "list":
				return b.sendSortedMarketsList(message, command.param[1])
			case "location", "country", "city", "info":
				p.page = 1
				p.query = command.param[0]
				p.queryData = concatenateStr(command.param[1:])

				return b.FindMarketsWithParams(message, &p)
			default:
				return b.unknownMessage(message)
			}
		default:
			return b.unknownMessage(message)
		}
	default:
		return b.unknownMessage(message)
	}
}

//Command start handler
func (b *Bot) startCommand(message *tgbotapi.Message) error {
	isUser, _ := b.storage.FindUser(uint(message.From.ID))
	if isUser.RowsAffected == 1 { //if find user
		return b.helloMessage(message)
	}

	switch message.From.LanguageCode {
	case ru:
		return b.createUser(message, b.config.RuDictionary)
	case en:
		return b.createUser(message, b.config.EnDictionary)
	default:
		return b.chooseLanguageMessage(message, b.config.EnDictionary)
	}
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
		return b.sendMessage(msg)
	case EnDictionary:
		if err := b.storage.CreateUser(&user); err != nil { //
			return err
		}

		msg := massegaConstructor(message, fmt.Sprintf("%s, %s", message.From.FirstName, d.startMessage))
		return b.sendMessage(msg)
	}
	return nil
}

func (b *Bot) sendSortedMarketsList(message *tgbotapi.Message, param string) error {
	res, err := b.storage.SortedMarketList(param)
	if err != nil {
		return err
	}

	switch param {
	case "markets":
		var m market

		m.market = res

		msg := massegaConstructor(message, *textParser(m))
		return b.sendMessage(msg)
	case "locations":
		var l location

		l.location = res

		msg := massegaConstructor(message, *textParser(l))
		return b.sendMessage(msg)
	case "countries":
		var c country

		c.country = res

		msg := massegaConstructor(message, *textParser(c))
		return b.sendMessage(msg)
	case "cities":
		var c city

		c.city = res

		msg := massegaConstructor(message, *textParser(c))
		return b.sendMessage(msg)
	}

	return nil
}
