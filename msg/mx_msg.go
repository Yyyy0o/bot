package msg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type MxMessage struct {
	Host     string
	Token    string
	lastTime float64
}

func (m *MxMessage) GetMessage() []Message {
	// 调用view接口
	viewBody := []byte(fmt.Sprintf(`{"rid":4617,"tt":%d}`, time.Now().Unix()))
	req, err := http.NewRequest("POST", m.Host+"/4/api/room/view", bytes.NewBuffer(viewBody))
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return nil
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("token", m.Token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("调用view出错...")
		return nil
	}
	defer resp.Body.Close()

	// 调用list接口
	listBody := []byte(fmt.Sprintf(`{"rid":4617,"msgid":0,"pagesize":30,"tt":%d}`, time.Now().Unix()))

	req, err = http.NewRequest("POST", m.Host+"/4/api/msg/list", bytes.NewBuffer(listBody))
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return nil
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("token", m.Token)

	client = &http.Client{}
	resp, err = client.Do(req)

	if err != nil {
		log.Println("获取消息列表出错...")
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

	if dataMap["code"] == float64(200) {
		if messages, ok := dataMap["list"].([]interface{}); ok {
			result := make([]Message, len(messages))
			current := m.lastTime
			for _, msgData := range messages {
				if msgData, ok := msgData.(map[string]interface{}); ok {
					if msgData["createtime"].(float64) > current {
						var msg []map[string]interface{}
						err := json.Unmarshal([]byte(msgData["msg"].(string)), &msg)
						if err != nil {
							fmt.Println("解析错误:", err)
							continue
						}

						switch msg[0]["type"].(string) {
						case "text":
							result = append(result, Message{
								Type:    "text",
								Content: msg[0]["msg"].(string),
							})
						case "pic":
							result = append(result, Message{
								Type:    "image",
								Content: msg[0]["url"].(string),
							})
						}

						m.lastTime = max(msgData["createtime"].(float64), m.lastTime)
					}
				}
			}
			return result
		}
	}

	return nil
}
