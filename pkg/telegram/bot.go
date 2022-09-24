package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	api *tgbotapi.BotAPI
}

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{api: bot}
}

func (b* Bot) Start() error {
	log.Printf("Authorized on account %s", b.api.Self.UserName)

	updates := b.initUpdatesChannel()

	b.handleUpdates(updates)

	return nil
}

func (b* Bot) initUpdatesChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.api.GetUpdatesChan(u)
}

func (b* Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf(update.Message.Text)
		log.Printf(update.CallbackData())

		if update.CallbackQuery != nil {
            // Respond to the callback query, telling Telegram to show the user
            // a message with the data received.
            callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
            if _, err := b.api.Request(callback); err != nil {
                panic(err)
            }

            // And finally, send a message containing the data received.
            msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
            if _, err := b.api.Send(msg); err != nil {
                panic(err)
            }
		}

		if update.Message.IsCommand() {
			b.handleCommand(update.Message)
			continue
		}

		b.handleMessage(update.Message)
	}
}