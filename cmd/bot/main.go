package main

import (
	"fmt"
	"log"
	"projects/staff-bot/pkg/models"
	"projects/staff-bot/pkg/telegram"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	models.ConnectDatabase()

	bot, err := tgbotapi.NewBotAPI("5408861618:AAFLxwtfNXnrelPPSEYqoOdjh1BnD2_tgSU")
	if err != nil {
		log.Panic(err)
	}
	
	telegramBot := telegram.NewBot(bot)

	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}

	fmt.Print("It's ok!")
}
