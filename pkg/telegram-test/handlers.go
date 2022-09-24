package telegramtest

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const commandStart = "start"

func (b *Bot) handeCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Unknown command")

	switch message.Command() {
	case commandStart:
		msg.Text = "Command /start"
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
