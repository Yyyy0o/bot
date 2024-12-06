package msg

import "testing"

func TestMxMessage_GetMessage(t *testing.T) {
	mx := &MxMessage{
		Host:  "",
		Token: "",
	}
	messages := mx.GetMessage()
	t.Logf("Message: %+v", messages)
}
