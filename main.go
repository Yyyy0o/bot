package main

import (
	"bot/wxbot"
)

func main() {
	b := wxbot.InitBot()
	wxbot.Run(b)
}
