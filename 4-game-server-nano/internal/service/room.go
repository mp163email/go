package service

import (
	"errors"
	localcomp "game-server-nano/internal/component" //包别名解决包名冲突
	"game-server-nano/internal/protocol"
	nanocomp "github.com/lonng/nano/component"
	"github.com/lonng/nano/session"
	"log"
)

// service->comps(存数据，操作数据)
// service是一个nona组件
type RoomService struct {
	manager *localcomp.RoomManager
	nanocomp.Base
	//匿名字段,不需要名字 1.嵌入nano框架组件,标记这是个服务组件 2.提供组件生命周期管理 3.自动过得nano路由能力
	//在调用的时候，外层可以直接调用 nanocomp.Base里的方法
}

func NewRoomService(mgr *localcomp.RoomManager) *RoomService {
	return &RoomService{
		manager: mgr,
	}
}

// 可以被nano路由 join
func (s *RoomService) Join(sin *session.Session, msg *protocol.ClientMessage) error {
	player := &protocol.Player{
		ID:      string(sin.UID()),
		Name:    "Player" + string(sin.UID()),
		Session: sin,
	}
	//客户端传来的消息 data可以转换成map
	data, ok := msg.Data.(map[string]interface{}) //.()是断言的意思，判断msg.Data是否是map[string]interface{}类型
	if !ok {
		return sin.Response(&protocol.ServerMessage{
			Type:    "join",
			Data:    "Invalid data format",
			Success: false,
			Error:   "Invalid data format",
		})
	}

	//处理能直接加入房间
	//要求data中必须有roomID字段，并且是string类型
	if roomID, ok := data["roomID"].(string); ok {
		_, err := s.manager.JoinRoom(roomID, player)
		if err != nil {
			return sin.Response(&protocol.ServerMessage{
				Type:    "join",
				Data:    "Failed to join room",
				Success: false,
				Error:   err.Error(),
			})
		}
		//广播新玩家加入
		s.Broadcast(roomID, "player_joined", player, string(sin.UID()))
		return sin.Response(&protocol.ServerMessage{
			Type:    "join",
			Data:    "Room joined",
			Success: true,
			Error:   "",
		})
	}

	//处理创建房间
	gameType, ok := data["gameType"].(string)
	if !ok {
		return errors.New("invalid game type") //errors.New返回的是一个error
	}
	room := s.manager.CreateRoom(gameType, player)
	log.Println("room created success~", room.ID)
	return sin.Response(&protocol.ServerMessage{
		Type:    "join",
		Data:    "Room created",
		Success: true,
		Error:   "",
	})
}

// 可以被nano路由
// 移动广播
func (s *RoomService) Move(sin *session.Session, msg *protocol.ClientMessage) error {
	room, err := s.manager.GetRoomByPlayerId(string(sin.UID()))
	if err != nil {
		log.Println("failed to find room:", err)
		return err
	}
	//广播
	s.Broadcast(room.ID, "game_move", msg.Data, string(sin.UID()))
	return nil
}

// 广播方法
func (s *RoomService) Broadcast(roomId, route string, data interface{}, exclude string) {
	room, err := s.manager.GetRoom(roomId)
	if err == nil {
		log.Println("failed to Broadcast:", err)
		return
	}
	for _, player := range room.Players {
		if player.ID != exclude {
			session := s.manager.GetSession(player.ID)
			if session != nil {
				session.Push(route, data)
			}
		}
	}
}
