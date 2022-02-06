package telegram

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"text/template"

	"github.com/Kodik77rus/fin-spy-tg-bot/internal/fin-spy-tg-bot/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var querys = [1]string{"all_markets"}

type Pagination struct {
	isValid bool
	query   string
	page    int
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

func paginationParser(params []string) *Pagination {
	pageNumber := strings.Split(params[0], "=")

	i, err := strconv.Atoi(pageNumber[1])
	if err != nil {
		return &Pagination{isValid: false}
	}

	query := strings.Split(params[1], "=")
	for _, q := range querys {
		if isGood := strings.Compare(q, query[1]); isGood == 0 {
			return &Pagination{
				isValid: true,
				query:   query[1],
				page:    i,
			}
		}
	}

	return &Pagination{isValid: false}
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
		txt := geoParser(i.location, "country")
		return txt
	case country:
		txt := geoParser(i.country, "country")
		return txt
	case city:
		txt := geoParser(i.city, "country")
		return txt
	}
	return nil
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

//Send default message for unknown command
func (b *Bot) unknownMessage(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Silly bot Finn don't understant you!")
	if _, err := b.bot.Send(msg); err != nil {
		panic(err)
	}
	return nil
}
