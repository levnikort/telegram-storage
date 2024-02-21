package bot

import (
	botApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/levnikort/telegram-storage/config"
)

type Bot struct {
	API *botApi.BotAPI
}

func New() *Bot {
	api, err := botApi.NewBotAPI(config.Config.TelegramBotToken)

	if err != nil {
		panic("fail connect bot")
	}

	return &Bot{
		API: api,
	}
}
