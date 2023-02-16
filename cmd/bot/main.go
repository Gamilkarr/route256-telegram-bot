package main

import (
	"log"

	"github.com/Gamilkarr/route256-telegram-bot/internal/clients/tg"
	"github.com/Gamilkarr/route256-telegram-bot/internal/config"
	"github.com/Gamilkarr/route256-telegram-bot/internal/model/messages"
)

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatal("config init failed:", err)
	}

	tgClient, err := tg.New(config)
	if err != nil {
		log.Fatal("tg client init failed:", err)
	}

	msgModel := messages.New(tgClient)

	tgClient.ListenUpdates(msgModel)
}
