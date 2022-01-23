package telegram

import (
	"bytes"
	"github.com/Kodik77rus/fin-spy-tg-bot/internal/fin-spy-tg-bot/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"text/template"
)

func massegaConstructor(message *tgbotapi.Message, text string) *tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	return &msg
}

func textParser(i interface{}) string {
	switch i := i.(type) {
	case models.Market:
		msg, err := marketParser(i)
		if err != nil {
			panic(err)
		}
		return msg
	}
	return ""
}

func marketParser(m models.Market) (string, error) {
	var buf bytes.Buffer

	ut, err := template.New("market").Parse(
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

func inlineKeyBoardConstructor(text string, data string) *tgbotapi.InlineKeyboardMarkup {
	switch text {
	case "info":
		keyBoard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonURL(text, data),
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
