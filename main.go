package main

import (
	"time"
	"wxbot/bot"
	"wxbot/msg"
)

func main() {
	bot := bot.NewBot()

	bot.Run()

	msgChan := make(chan string)

	go getQQMessage(msgChan)

	for {
		for message := range msgChan {
			bot.SendMessage("yyyyy", message)
		}
	}
}

func getQQMessage(msgChan chan string) {
	for {
		msg.QQMessage(msgChan)
		time.Sleep(1 * time.Minute)
	}
}
