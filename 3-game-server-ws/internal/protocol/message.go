package protocol

import (
	"strconv"
	"time"
)

// 后面用json说明这个类可能会和json做交互 即从json串转到类  从类转到json串
// 定义一个消息类
type Message struct {
	ID   MsgID                  `json:"id"`
	Data map[string]interface{} `json:"data"`
	Time time.Time              `json:"time"`
}

// 聊天消息
type ChatMessage struct {
	From    string `json:"from"`
	Content string `json:"content"`
}

// 错误消息
type ErrorMessage struct {
	Reason string `json:"reason"`
}

// 创建一个新的消息对象
func NewMessage(id MsgID, data interface{}) Message {
	idStr := strconv.Itoa(int(id))
	return Message{
		ID:   id,
		Data: toMap(idStr, data),
		Time: time.Now(),
	}
}

// 创建一个心跳消息对象
func NewHeartbeatMessage() Message {
	return NewMessage(MsgHeartbeat, nil)
}

// 创建一个聊天消息对象
func NewChatMessage(from, content string) Message {
	return NewMessage(MsgChat, ChatMessage{from, content})
}

// 创建一个错误消息对象
func NewErrorMessage(reason string) Message {
	return NewMessage(MsgError, reason)
}

// 创建一个系统消息对象
func NewSystemMessage(content string) Message {
	return NewMessage(MsgSystem, content)
}

func toMap(key string, data interface{}) map[string]interface{} {
	mp := make(map[string]interface{})
	mp[key] = data
	return mp
}
