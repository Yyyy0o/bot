package wxbot

import (
	"log"

	"github.com/eatmoreapple/openwechat"
)

func messageHandler(msg *openwechat.Message) {
	log.Printf("收到消息 : %s", msg.Content)

}
