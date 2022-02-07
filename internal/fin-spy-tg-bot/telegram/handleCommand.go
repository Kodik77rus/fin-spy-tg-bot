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
			var pagination Pagination

			markets, _ := b.storage.GetAllMarkets(1) //firts page

			for _, m := range markets.Markets {
				msg := massegaConstructor(message, *textParser(m))
				msg.ReplyMarkup = inlineKeyBoardConstructor("info", m.Hour)
				b.sendMessage(msg)
			}

			pagination.page = 1
			pagination.query = "all_markets"

			return b.paginationMessage(message, &pagination)
		case "location", "country", "city":
			var p Pagination

			p.page = 1
			p.query = flags[1]
			p.queryData = parseQuery(flags[2:])

			markets, _ := b.storage.FindMarketsWithGeoParams(p.query, p.queryData, 1)
			if markets.Count == 0 {
				msg := massegaConstructor(message, "Markets not found")
				b.sendMessage(msg)
				return nil
			}

			for _, m := range markets.Markets {
				msg := massegaConstructor(message, *textParser(m))
				msg.ReplyMarkup = inlineKeyBoardConstructor("info", m.Hour)
				b.sendMessage(msg)
			}

			if markets.Count == 1 {
				return nil
			}

			return b.paginationMessage(message, &p)
		default:
			return b.unknownMessage(message)
		}
	default:
		return b.unknownMessage(message)
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

func (b *Bot) findMarketsWithParam(message *tgbotapi.Message, param string) error {
	res, err := b.storage.FindMarketsWithParam(param)
	if err != nil {
		return err
	}

	switch param {
	case "location":
		var location location

		location.location = res

		msg := massegaConstructor(message, *textParser(location))
		return b.sendMessage(msg)
	case "country":
		var country country

		country.country = res

		msg := massegaConstructor(message, *textParser(country))
		return b.sendMessage(msg)
	case "city":
		var city city

		city.city = res

		msg := massegaConstructor(message, *textParser(city))
		return b.sendMessage(msg)
	default:
		b.unknownMessage(message)
	}
	return nil
}
