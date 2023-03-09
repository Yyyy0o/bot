package bot

import (
	"log"

	"github.com/eatmoreapple/openwechat"
)

type WechatBot struct {
	bot *openwechat.Bot

	botName string
}

var (
	NORMAL  = "normal"
	DESKTOP = "desktop"
)

func (w *WechatBot) Init(cfg *cfg.Config) error {
	switch cfg.BotConfig.WechatLoginType {
	case DESKTOP:
		w.bot = openwechat.DefaultBot(openwechat.Desktop)
	case NORMAL:
		w.bot = openwechat.DefaultBot(openwechat.Normal)
	default:
		w.bot = openwechat.DefaultBot(openwechat.Normal)
	}

	log.Printf("[Init] wechat bot init success, bot name: %s", w.botName)
	return nil
}
