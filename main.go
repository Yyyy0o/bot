package main

import (
	"bot/bot"
	"bot/msg"
	"bot/util"
	"fmt"
	"time"
)

var globalConfig util.Config
var globalBot *bot.Bot
var messageChan chan msg.Message

func main() {
	initBot()
	start()

	for {
		for message := range messageChan {
			globalBot.SendMessage("yyyyy", message)
		}
	}
}

func initBot() {
	var err error
	globalConfig, err = util.LoadConfig("config.yaml")
	if err != nil {
		fmt.Printf("load config error: %+v", err.Error())
		return
	}

	globalBot = bot.NewBot()
	globalBot.Run()

	messageChan = make(chan msg.Message)
}

func start() {
	mx := &msg.MxMessage{
		Host:  globalConfig.MX.Host,
		Token: globalConfig.MX.Token,
	}

	qq := &msg.QQMessage{
		Host:  globalConfig.QQ.Host,
		Group: globalConfig.QQ.Group,
	}

	go getMessage(mx)
	go getMessage(qq)
}

func getMessage(mp msg.MessageProducer) {
	for {
		messages := mp.GetMessage()
		for _, msg := range messages {
			messageChan <- msg
		}
		time.Sleep(1 * time.Minute)
	}
}
