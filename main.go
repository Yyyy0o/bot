package main

import (
	"bot/config"
	"bot/wxbot"
)

func main() {
	cfg, err := config.InitConfig()

	if err != nil {
		panic(err)
	}

	b := wxbot.InitBot(cfg)
	wxbot.Run(b)
}
