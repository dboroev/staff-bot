package telegramtest

import (
	"log"
	"projects/staff-bot/pkg/repository"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	api             *tgbotapi.BotAPI
	tokenRepostiory repository.TokenRepository
}

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{api: bot}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.api.Self.UserName)

	updates := b.initUpdatesChannel()

	b.handeUpdates(updates)

	return nil
}

func (b *Bot) handeUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			b.handeCommand(update.Message)
			continue
		}

		b.handleMessage(update.Message)
	}
}

func (b *Bot) initUpdatesChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.api.GetUpdatesChan(u)
}
