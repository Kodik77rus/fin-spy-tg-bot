package telegram

import (
	"github.com/Kodik77rus/fin-spy-tg-bot/internal/fin-spy-tg-bot/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var inlineLanguageKeyBoard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(ru, ru),
		tgbotapi.NewInlineKeyboardButtonData(en, en),
	),
)

//fimd user in db and change user language
func (b *Bot) setUserLanguage(user *models.User) error {
	if err := b.storage.UpdateUser(user); err != nil {
		return err
	}
	return  nil
}

func massegaConstructor(message *tgbotapi.Message, text string) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(message.Chat.ID, text)
}
