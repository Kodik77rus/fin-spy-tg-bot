package telegram

import (
	"fmt"

	"github.com/Kodik77rus/fin-spy-tg-bot/internal/fin-spy-tg-bot/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	commandStart = "start"
	// commandWhatch     = "exchange"
	// commandWhatch     = "whatch"
	// commandDelete     = "delete"
	// commandWhatchList = "whatchlist"
	// commandInfo       = "info"

	ru = "ru"
	en = "en"
)

var inlineLanguageKeyBoard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(ru, ru),
		tgbotapi.NewInlineKeyboardButtonData(en, en),
	),
)

var user models.User

//Handle commands
func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	//message constructor
	baseMsg := tgbotapi.NewMessage(message.Chat.ID, "base message")

	switch message.Command() {
	case commandStart:
		msg, err := b.startCommand(message, &baseMsg)
		if err != nil {
			return err
		}
		baseMsg = *msg

	default:
		return b.unknownCommand(message)
	}

	if _, err := b.bot.Send(baseMsg); err != nil {
		return err
	}
	return nil
}

//Handle callback querys
func (b *Bot) callbackQueryHandler(cb *tgbotapi.CallbackQuery) error {
	msg := tgbotapi.NewMessage(cb.Message.Chat.ID, "base message") // "base message" is crutch

	// Respond to the callback query, telling Telegram to show the user
	// a message with the data received.
	callback := tgbotapi.NewCallback(cb.ID, cb.Data)
	if _, err := b.bot.Request(callback); err != nil {
		panic(err)
	}

	user.Id = uint(cb.Message.Chat.ID)
	user.Language = cb.Data

	switch cb.Data {
	case ru:
		str, err := b.setUserLanguage(&user)
		if err != nil {
			return err
		}
		msg.Text = str
	case en:
		str, err := b.setUserLanguage(&user)
		if err != nil {
			return err
		}
		msg.Text = str
	}

	//Send a message containing the data received.
	if _, err := b.bot.Send(msg); err != nil {
		panic(err)
	}

	return nil
}

//Send default message for unknown command
func (b *Bot) unknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Silly bot Finn don't understant you!")
	if _, err := b.bot.Send(msg); err != nil {
		panic(err)
	}
	return nil
}

//Command start handler
func (b *Bot) startCommand(message *tgbotapi.Message, msg *tgbotapi.MessageConfig) (*tgbotapi.MessageConfig, error) {
	isUser, err := b.storage.FindUser(uint(message.From.ID))
	if err != nil {
		msg.Text = "troubls with bd"
		return msg, err
	}

	//if find user
	if isUser.RowsAffected == 1 {
		msg.Text = fmt.Sprintf("Hello %s!", message.From.FirstName)
		return msg, nil
	}

	user.Id = uint(message.From.ID)
	user.UserName = message.From.FirstName
	user.Language = message.From.LanguageCode

	switch message.From.LanguageCode {
	case ru:
		if err := b.storage.CreateUser(&user); err != nil {
			return msg, err
		}

		msg.Text = "RU"
		return msg, nil
	case en:
		if err := b.storage.CreateUser(&user); err != nil {
			return msg, err
		}

		msg.Text = "EN"
		return msg, nil
	default:
		msg.ReplyMarkup = inlineLanguageKeyBoard
		msg.Text = "Choose language"
	}

	return msg, nil
}

//fimd user in db and change user language
func (b *Bot) setUserLanguage(user *models.User) (string, error) {
	if err := b.storage.UpdateUser(user); err != nil {
		return "", err
	}
	return "you chose language", nil
}

// func inlineKeyBoardConstructor func () tgbotapi.InlineKeyboardMarkup {

// }
