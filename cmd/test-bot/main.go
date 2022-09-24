package main

import (
	"log"
	"projects/staff-bot/pkg/config"
	"projects/staff-bot/pkg/telegram"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	telegramBot := telegram.NewBot(bot)

	// go func ()  {
	// 	if err := telegramBot.Start(); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	// authorizationServer = server.NewAutoriztionServer()

	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}
