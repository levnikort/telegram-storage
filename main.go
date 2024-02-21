package main

import (
	"github.com/bluele/gcache"
	"github.com/go-martini/martini"
	"github.com/levnikort/telegram-storage/bot"
	"github.com/levnikort/telegram-storage/config"
)

func main() {
	config.Load()

	controller := Controller{
		bot.New(),
		gcache.New(config.Config.CacheElements).
			Expiration(config.Config.CacheExpirationDate).
			Build(),
	}
	s := martini.Classic()

	s.Get("/:file_id", controller.Download)
	s.Post("/upload/:file_type", controller.Upload)

	s.RunOnAddr(":" + config.Config.HttpServerPort)
}
