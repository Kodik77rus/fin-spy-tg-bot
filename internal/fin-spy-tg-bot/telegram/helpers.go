package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var inlineLanguageKeyBoard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(ru, ru),
		tgbotapi.NewInlineKeyboardButtonData(en, en),
	),
)

func massegaConstructor(message *tgbotapi.Message, text string) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(message.Chat.ID, text)
}
