package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	TelegramBotToken    string
	TelegramChatID      int
	HttpServerPort      string
	CacheExpirationDate time.Duration
	CacheElements       int
}

var Config *AppConfig = new(AppConfig)

func Load() {
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}

	value, ok := os.LookupEnv("TELEGRAM_BOT_TOKEN")

	if !ok {
		panic("No TELEGRAM_BOT_TOKEN variable")
	}

	Config.TelegramBotToken = value

	value, ok = os.LookupEnv("HTTP_SERVER_PORT")

	if !ok {
		panic("No HTTP_SERVER_PORT variable")
	}

	Config.HttpServerPort = value

	value, ok = os.LookupEnv("TELEGRAM_CHAT_ID")

	if !ok {
		panic("No TELEGRAM_CHAT_ID variable")
	}

	id, err := strconv.Atoi(value)

	if err != nil {
		panic("No TELEGRAM_CHAT_ID convert string to int")
	}

	Config.TelegramChatID = id

	value, ok = os.LookupEnv("CACHE_EXPIRATION_DATE")

	if !ok {
		panic("No CACHE_EXPIRATION_DATE variable")
	}

	num, err := strconv.Atoi(value)

	if err != nil {
		panic("No CACHE_EXPIRATION_DATE convert string to int")
	}

	Config.CacheExpirationDate = time.Duration(num) * time.Minute

	value, ok = os.LookupEnv("CACHE_ELEMENTS")

	if !ok {
		panic("No CACHE_ELEMENTS variable")
	}

	num, err = strconv.Atoi(value)

	if err != nil {
		panic("No CACHE_ELEMENTS convert string to int")
	}

	Config.CacheElements = num
}
