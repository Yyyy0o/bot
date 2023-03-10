package wxbot

import (
	"fmt"
	"log"

	"github.com/eatmoreapple/openwechat"
	qrcode "github.com/skip2/go-qrcode"
)

type WechatBot struct {
	bot *openwechat.Bot
}

func InitBot() *WechatBot {
	b := &WechatBot{}

	cfg, err := initConfig()

	if err != nil {
		panic(err)
	}

	err = b.init(cfg)
	if err != nil {
		panic(err)
	}

	return b
}

func (w *WechatBot) init(cfg *Config) error {
	if cfg.LoginType == "desktop" {
		w.bot = openwechat.DefaultBot(openwechat.Desktop)
	} else {
		w.bot = openwechat.DefaultBot(openwechat.Normal)
	}

	// 打印登录二维码
	w.bot.UUIDCallback = handleQrCode
	// 消息处理
	w.bot.MessageHandler = messageHandler

	log.Printf("[Init] wechat bot init success")
	return nil
}

func Run(w *WechatBot) {
	// 免扫码登录
	reloadStorage := openwechat.NewFileHotReloadStorage("storage.json")
	defer reloadStorage.Close()
	err := w.bot.PushLogin(reloadStorage, openwechat.NewRetryLoginOption())

	if err != nil {
		panic(err)
	}

	log.Println("[Run] wechat bot is running")

	//阻塞主程序
	w.bot.Block()
}

func handleQrCode(uuid string) {
	q, _ := qrcode.New("https://login.weixin.qq.com/l/"+uuid, qrcode.Low)
	fmt.Println(q.ToString(true))
}
