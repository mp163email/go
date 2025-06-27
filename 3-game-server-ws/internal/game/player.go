package game

import (
	"game-server-ws/internal/config"
	"game-server-ws/internal/protocol"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"time"
)

type Player struct {
	ID               string
	Conn             *websocket.Conn
	lastHeartbeat    time.Time
	heartbeatTimeout time.Duration
}

func NewPlayer(conn *websocket.Conn) *Player {

	return &Player{
		ID:               generateID(),
		Conn:             conn,
		lastHeartbeat:    time.Now(),
		heartbeatTimeout: config.GlobalConfig.Game.HeartbeanInterval, //30秒无心跳断开连接
	}
}

// 发送消息方法
func (p *Player) Send(msg protocol.Message) error {
	return p.Conn.WriteJSON(msg)
}

// 处理心跳方法
func (p *Player) StartHeartbeat() {
	ticker := time.NewTicker(5 * time.Second)
	go func() { //异步监听ticker
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C: //每5秒会向ticker.C通道（channel)发数据
				if time.Since(p.lastHeartbeat) > p.heartbeatTimeout { //计算距离上次心跳的时间，如果超过30秒，断开连接
					p.Conn.Close()
					return
				}
				p.Send(protocol.NewHeartbeatMessage())
			}
		}
	}()
}

// 更新最后心跳时间方法
func (p *Player) UpdateLastHeartbeat() {
	p.lastHeartbeat = time.Now()
}

func generateID() string {
	// 这里可以使用 UUID 或其他方式生成唯一 ID
	id := uuid.New().String()
	return id
}
