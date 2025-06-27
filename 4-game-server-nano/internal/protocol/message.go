package protocol

import "github.com/lonng/nano/session"

// 客户端消息类型
const (
	MsgTypeLogin     = "login"
	MsgTypeJoin      = "join"
	MsgTypeLeave     = "leave"
	MsgTypeMove      = "move"
	MsgTypeChat      = "chat"
	MsgTypeHeartbeat = "heartbeat"
)

// 客户端发消息结构
type ClientMessage struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

// 服务器发消息结构
type ServerMessage struct {
	Type    string `json:"type"`
	Data    string `json:"data"`
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

// 玩家结构
type Player struct {
	ID      string           `json:"id"`
	Name    string           `json:"name"`
	Session *session.Session `json:"-"` //-表示不序列化这个字段
}

// 房间结构
type Room struct {
	ID       string    `json:"id"`
	Players  []*Player `json:"players"`  // 房间中的玩家列表
	State    int       `json:"state"`    // 房间状态：0-等待中，1-游戏中，2-已结束
	GameType string    `json:"gameType"` //什么类型的游戏， 棋牌， 麻将
}
