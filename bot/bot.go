package bot

import (
	"bot/msg"
	"fmt"
	"os"

	"github.com/eatmoreapple/openwechat"
	"github.com/skip2/go-qrcode"
)

type Bot struct {
	wxbot *openwechat.Bot
	Alive bool
}

func NewBot() *Bot {
	bot := &Bot{
		Alive: false,
		wxbot: openwechat.DefaultBot(openwechat.Desktop),
	}

	// 拼接桌面路径
	path := os.Getenv("USERPROFILE") + "/Desktop/qrcode.png"

	bot.wxbot.UUIDCallback = func(uuid string) {
		q, _ := qrcode.New("https://login.weixin.qq.com/l/"+uuid, qrcode.Low)
		q.WriteFile(400, path)
		fmt.Printf("扫描%s路径下二维码登录\n", path)
	}

	bot.wxbot.LoginCallBack = func(body openwechat.CheckLoginResponse) {
		bot.Alive = true
	}

	return bot
}

func (b *Bot) Run() {
	// 登录
	reloadStorage := openwechat.NewFileHotReloadStorage("storage.json")
	defer reloadStorage.Close()
	err := b.wxbot.PushLogin(reloadStorage, openwechat.NewRetryLoginOption())
	if err != nil {
		b.Alive = false
		fmt.Printf("登录失败... %v\n", err)
	}
}

func (b *Bot) SendMessage(name string, message msg.Message) {
	if message.Content == "" {
		return
	}
	fmt.Printf("bot %t 发送消息%s : %s\n", b.Alive, name, message)
	if !b.Alive {
		return
	}
	self, err := b.wxbot.GetCurrentUser()
	if err != nil {
		println("获取当前用户失败")
		return
	}
	friends, err := self.Friends()
	if err != nil {
		println("获取好友列表失败")
		return
	}
	for _, friend := range friends {
		if friend.NickName == name {
			friend.SendText(message.Content)
			return
		}
	}
}
