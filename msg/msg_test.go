package msg

import "testing"

func TestMxMessage_GetMessage(t *testing.T) {
	mx := &MxMessage{
		Host:  "https://mx.tg0536.cn",
		Token: "1c6db64ad048ec84fb586929e67c42f9",
	}
	messages := mx.GetMessage()
	t.Logf("Message: %+v", messages)
}
