package game

import "net"

type Player struct {
	ID   string   // 玩家ID
	Conn net.Conn // 玩家连接
}
