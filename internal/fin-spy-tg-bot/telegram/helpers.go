package telegram

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"text/template"

	"github.com/Kodik77rus/fin-spy-tg-bot/internal/fin-spy-tg-bot/models"
	"github.com/Kodik77rus/fin-spy-tg-bot/storage"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var querys = [4]string{"all_markets", "location", "country", "city"}

type Pagination struct {
	isValid   bool
	query     string
	queryData string
	page      int
}

type location struct {
	location *[]string
}

type country struct {
	country *[]string
}

type city struct {
	city *[]string
}

func textParser(i interface{}) *string {
	switch i := i.(type) {
	case models.Market:
		txt, err := marketParser(i)
		if err != nil {
			panic(err)
		}
		return &txt
	case location:
		txt := geoParser(i.location, "location")
		return txt
	case country:
		txt := geoParser(i.country, "country")
		return txt
	case city:
		txt := geoParser(i.city, "city")
		return txt
	}
	return nil
}

func paginationParser(params []string) *Pagination {
	pageNumber := strings.Split(params[0], "=")

	i, err := strconv.Atoi(pageNumber[1])
	if err != nil {
		return &Pagination{isValid: false}
	}

	query := strings.Split(params[1], "=")
	queryData := strings.Split(params[2], "=")

	for _, q := range querys {
		if isGood := strings.Compare(q, query[1]); isGood == 0 {
			return &Pagination{
				isValid:   true,
				query:     q,
				queryData: queryData[1],
				page:      i,
			}
		}
	}

	return &Pagination{isValid: false}
}

func marketParser(m models.Market) (string, error) {
	var buf bytes.Buffer

	ut, err := template.New("market").
		Parse(
			"Name: {{ .Name }}\n" +
				"Code: {{ .Code }}\n" +
				"Mic: {{ .Mic }}\n" +
				"Location: {{ .Location }}\n" +
				"Country: {{ .Country }}\n" +
				"City: {{ .City }}\n" +
				"Delay: {{ .Delay }} min\n",
		)
	if err != nil {
		return "", err
	}

	err = ut.Execute(&buf, m)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func geoParser(location *[]string, param string) *string {
	str := new(string)
	for _, c := range *location {
		*str += fmt.Sprintf("%s:\n/markets_show_%s_%s\n", c, param, c)
	}
	return str
}

func massegaConstructor(message *tgbotapi.Message, text string) *tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	return &msg
}

func parseQuery(query []string) string {
	if len(query) == 1 {
		return query[0]
	}

	if len(query) == 2 {
		return query[0] + "_" + query[1]
	}
	return ""
}

func inlineKeyBoardConstructor(text string, data string) *tgbotapi.InlineKeyboardMarkup {
	param := strings.Split(text, " ")
	switch param[0] {
	case "info":
		keyBoard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonURL(text, data),
			),
		)
		return &keyBoard
	case "next":
		keyBoard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(strings.Title(text), data),
			),
		)
		return &keyBoard
	default:
		keyBoard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(ru, ru),
				tgbotapi.NewInlineKeyboardButtonData(en, en),
			),
		)
		return &keyBoard
	}
}

func (b *Bot) sendMarkets(message *tgbotapi.Message, markets *storage.MarketResponse) error {
	for _, m := range markets.Markets {
		msg := massegaConstructor(message, *textParser(m))
		msg.ReplyMarkup = inlineKeyBoardConstructor("info", m.Hour)
		b.sendMessage(msg)
	}
	return nil
}

//Send default message for unknown command
func (b *Bot) unknownMessage(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Silly bot Finn don't understant you!")
	return b.sendMessage(&msg)
}

func (b *Bot) paginationMessage(message *tgbotapi.Message, p *Pagination) error {
	msg := massegaConstructor(message, "Touch to see next markets")
	msg.ReplyMarkup = inlineKeyBoardConstructor(
		fmt.Sprintf("next page %d", p.page+1),
		fmt.Sprintf("page=%d,query=%s,queryData=%s", p.page+1, p.query, p.queryData),
	)
	return b.sendMessage(msg)
}

func (b *Bot) chooseLanguageMessage(message *tgbotapi.Message, dictionary interface{}) error {
	msg := massegaConstructor(message, "Choose language")
	msg.ReplyMarkup = inlineKeyBoardConstructor("", "") //crutch
	return b.sendMessage(msg)
}

func (b *Bot) helloMessage(message *tgbotapi.Message) error {
	msg := massegaConstructor(message, fmt.Sprintf("Hello %s!", message.From.FirstName))
	return b.sendMessage(msg)
}

func (b *Bot) sendMessage(msg *tgbotapi.MessageConfig) error {
	if _, err := b.bot.Send(msg); err != nil {
		panic(err)
	}
	return nil
}
