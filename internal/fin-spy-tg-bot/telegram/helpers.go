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

var querys = [5]string{"market", "location", "country", "city", "info"}

type pagination struct {
	isValid   bool
	query     string
	queryData string
	page      int
}

type command struct {
	name  string
	flag  string
	param []string
}

type market struct {
	market *[]string
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
	case market:
		return marketListParser(i.market, "info")
	case location:
		return marketListParser(i.location, "location")
	case country:
		return marketListParser(i.country, "country")
	case city:
		return marketListParser(i.city, "city")
	}
	return nil
}

func paginationParser(params []string) *pagination {
	pageNumber := strings.Split(params[0], "=")

	i, err := strconv.Atoi(pageNumber[1])
	if err != nil {
		return &pagination{isValid: false}
	}

	query := strings.Split(params[1], "=")
	queryData := strings.Split(params[2], "=")

	for _, q := range querys {
		if isGood := strings.Compare(q, query[1]); isGood == 0 {
			return &pagination{
				isValid:   true,
				query:     q,
				queryData: queryData[1],
				page:      i,
			}
		}
	}

	return &pagination{isValid: false}
}

func marketParser(m models.Market) (string, error) {
	var buf bytes.Buffer

	m.Name = replaceStr(m.Name)
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

func marketListParser(location *[]string, param string) *string {
	str := new(string)
	for _, c := range *location {
		*str += fmt.Sprintf("%s:\n/market_show_%s_%s\n", replaceStr(c), param, c)
	}
	return str
}

func commandValidation(c []string) *command {
	return &command{
		name:  c[0],
		flag:  c[1],
		param: c[2:],
	}
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

func (b *Bot) paginationMessage(message *tgbotapi.Message, p *pagination) error {
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

func (b *Bot) sendAllMarkets(message *tgbotapi.Message, page int) error {
	markets, err := b.storage.GetAllMarkets(page)
	if err != nil {
		return err
	}

	if markets.Count == 0 {
		msg := massegaConstructor(message, "You watched all markets!")
		return b.sendMessage(msg)
	}

	var p pagination

	b.sendMarkets(message, markets)

	p.page = page
	p.query = "market"
	p.queryData = "all"
	return b.paginationMessage(message, &p)
}

func (b *Bot) FindMarketsWithParams(message *tgbotapi.Message, p *pagination) error {
	markets, _ := b.storage.FindMarketsWithParams(p.query, p.queryData, p.page)
	fmt.Println(markets.Count)
	if markets.Count == 0 {
		msg := massegaConstructor(message, "you see all markets")
		b.sendMessage(msg)
		return nil
	}

	b.sendMarkets(message, markets)

	if markets.Count == 1 {
		return nil
	}

	return b.paginationMessage(message, p)
}

func replaceStr(s string) string {
	return strings.ReplaceAll(s, "_", " ")
}

func concatenateStr(s []string) string {
	var str strings.Builder

	arrLen := len(s)

	for i, e := range s {
		if i == arrLen-1 {
			str.WriteString(e)
			continue
		}
		str.WriteString(e + "_")
	}

	return str.String()
}
