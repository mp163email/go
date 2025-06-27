package component

import (
	"errors"
	"game-server-nano/internal/protocol"
	"github.com/google/uuid"
	"github.com/lonng/nano/session"
	"sync"
)

// 定义几个变量
var (
	ErrRoomNotFound    = errors.New("room not found") //使用errors.New()创建一个错误对象
	ErrRoomFull        = errors.New("room is full")
	ErrPlayerNotInRoom = errors.New("player not in room")
	ErrInvalidPlayer   = errors.New("invalid player")
)

// 定义几个常量
const (
	GameTypeChess   = "chess"
	GameTypePoker   = "poker"
	MaxChessPlayers = 2
	MaxPokerPlayers = 4
)

// 房间管理组件
type RoomManager struct {
	rooms       map[string]*protocol.Room
	playerReals map[string]*protocol.Player //key:playerID value:player
	players     map[string]string           //key:playerID value:roomId
	rwLock      sync.RWMutex                //一个读写锁 读音：缪-泰克斯
}

// new一个房间管理组件
func NewRoommanager() *RoomManager {
	return &RoomManager{
		rooms:   make(map[string]*protocol.Room),
		players: make(map[string]string),
	}
}

func (m *RoomManager) GetSession(playerId string) *session.Session {
	m.rwLock.RLock()
	defer m.rwLock.RUnlock()
	player, ok := m.playerReals[playerId]
	if !ok {
		return nil
	}
	return player.Session
}

// 获取房间
func (m *RoomManager) GetRoom(roomId string) (*protocol.Room, error) {
	m.rwLock.RLock()
	defer m.rwLock.RUnlock()
	room, ok := m.rooms[roomId]
	if !ok {
		return nil, ErrRoomNotFound
	}
	return room, nil
}

// 根据玩家id获取房间
func (m *RoomManager) GetRoomByPlayerId(playerId string) (*protocol.Room, error) {
	m.rwLock.RLock()
	defer m.rwLock.RUnlock()
	roomId, ok := m.players[playerId]
	if !ok {
		return nil, ErrPlayerNotInRoom
	}
	room, ok := m.rooms[roomId]
	if !ok {
		return nil, ErrRoomNotFound
	}
	return room, nil
}

// 创建房间
// *代表是指针类型， 指针是引用类型， 可以修改值。 当需要修改值或者结构体较大时（直接指向地址,避免拷贝开销），使用指针。
func (m *RoomManager) CreateRoom(gameType string, creator *protocol.Player) *protocol.Room {
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	roomId := generateRoomId()
	room := protocol.Room{ //构造方法赋值
		ID:       roomId,
		GameType: gameType,
		State:    0,
		Players:  []*protocol.Player{creator},
	}
	m.rooms[roomId] = &room
	m.players[creator.ID] = roomId
	m.playerReals[creator.ID] = creator
	return &room
}

// 加入房间
func (m *RoomManager) JoinRoom(roomId string, player *protocol.Player) (*protocol.Room, error) {
	m.rwLock.RLock()
	defer m.rwLock.RUnlock()

	//检查玩家是否在其他房间
	if _, ok := m.players[player.ID]; ok {
		return nil, ErrInvalidPlayer
	}

	//检查房间是否存在
	room, ok := m.rooms[roomId]
	if !ok {
		return nil, ErrRoomNotFound
	}

	//检查房间人数上限
	maxPlayers := MaxChessPlayers
	if room.GameType == GameTypePoker {
		maxPlayers = MaxPokerPlayers
	}
	if len(room.Players) >= maxPlayers {
		return nil, ErrRoomFull
	}

	room.Players = append(room.Players, player) //操作room
	m.players[player.ID] = roomId               //操作manager
	m.playerReals[player.ID] = player
	return room, nil
}

// 离开房间
func (m *RoomManager) LeaveRoom(playerId string) (*protocol.Room, error) {
	m.rwLock.Lock()
	defer m.rwLock.Unlock()

	//检查玩家是否在房间内
	roomId, ok := m.players[playerId]
	if !ok {
		return nil, ErrRoomNotFound
	}

	//检查房间是否存在
	room, ok := m.rooms[roomId]
	if !ok {
		return nil, ErrRoomNotFound
	}

	//从房间中移除玩家
	for i, player := range room.Players {
		if player.ID == playerId {
			room.Players = append(room.Players[:i], room.Players[i+1:]...) //展开操作符,将切片中的元素展开， 因为append的第2个参数是一个一个的元素，不是一个切片，所以要把切片展开
			break
		}
	}

	delete(m.players, playerId)
	delete(m.playerReals, playerId)

	//房间内没人，删除房间
	if len(room.Players) == 0 {
		delete(m.rooms, roomId)
	}

	return room, nil

}

func generateRoomId() string {
	return uuid.New().String()
}
