package msg

import "testing"

func TestViewReq(t *testing.T) {
	result := viewReq()
	t.Logf("Message: %+v", result)
}

func TestQueryMsg(t *testing.T) {
	result := queryMsg()
	t.Logf("Message: %+v", result)
}

func TestMxMessage(t *testing.T) {
	msgChan := make(chan string)

	go MxMessage(msgChan)

	for msg := range msgChan {
		t.Logf("Message: %+v", msg)
	}
}
