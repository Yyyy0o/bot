package msg

import (
	"testing"
)

func TestStart(t *testing.T) {
	msgChan := make(chan string)

	go QQMessage(msgChan)

	for msg := range msgChan {
		t.Logf("Message: %+v", msg)
	}
}

func TestGetUrl(t *testing.T) {
	ary := getUrl()
	t.Logf("Message: %+v", ary)
}
