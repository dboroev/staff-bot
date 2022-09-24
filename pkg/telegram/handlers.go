package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const commandStart = "start"
const commandLogin = "login"

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Unknown command")

	switch message.Command() {
	case commandStart:
		msg.Text = "Приветствую в стафф-бот!"
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonURL("1.com", "http://1.com"),
				tgbotapi.NewInlineKeyboardButtonData("2", "2"),
				tgbotapi.NewInlineKeyboardButtonData("3", "3"),
			),
		)
	case commandLogin:
		msg.Text = "i love you"
		msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("1"),
				tgbotapi.NewKeyboardButton("2"),
			),
		)
	}

	_, err := b.api.Send(msg)
	
	return err
}

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	msg.ReplyToMessageID = message.MessageID

	b.api.Send(msg)
}