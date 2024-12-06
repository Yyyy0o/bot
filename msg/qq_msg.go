package msg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type QQMessage struct {
	Host     string
	Group    string
	lastTime float64
}

func (q *QQMessage) GetMessage() []Message {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetPrefix("[QQ_MSG]")

	reqBody := []byte(fmt.Sprintf(`{"group_id": %s,"count": 10,"reverseOrder": true}`, q.Group))

	resp, err := http.Post(q.Host+"/get_group_msg_history", "application/json", bytes.NewBuffer(reqBody))

	if err != nil {
		log.Println("获取消息出错...")
		return nil
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("读取消息出错...")
		return nil
	}

	var dataMap map[string]interface{}

	err = json.Unmarshal([]byte(body), &dataMap)
	if err != nil {
		log.Println("解析消息出错...")
		return nil
	}

	if dataMap["status"] == "ok" {
		if data, ok := dataMap["data"].(map[string]interface{}); ok {
			if messages, ok := data["messages"].([]interface{}); ok {
				result := make([]Message, len(messages))
				current := q.lastTime
				for i, v := range messages {
					if msg, ok := v.(map[string]interface{}); ok {
						time := msg["time"].(float64)
						if time > current {
							result[i] = Message{
								Content: msg["raw_message"].(string),
								Type:    "text",
							}
							q.lastTime = max(q.lastTime, time)
						}
					}
				}
				return result
			}
		}
	}

	return nil
}
