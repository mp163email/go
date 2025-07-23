package game

import (
	"game-server-ws/internal/config"
	"game-server-ws/internal/protocol"
	"sync"
)

/**
当前代码使用了 sync.RWMutex 读写锁，这能有效避免并发读写 map 的问题，其工作原理如下：
Join 和 Leave 方法（写操作）
Join 和 Leave 方法在修改 players map 时使用了写锁（Lock 和 Unlock）：
写锁是独占锁，当一个 goroutine 获取写锁后，other goroutine 既不能获取写锁，也不能获取读锁，直到写锁被释放。这意味着在 Join 或 Leave 方法执行期间，Broadcast 方法无法对 players 进行遍历。

Broadcast 方法（读操作）
Broadcast 方法在遍历 players map 时使用了读锁（RLock 和 RUnlock）：
读锁允许多个 goroutine 同时获取，也就是说多个 Broadcast 方法可以并发执行。但在有 goroutine 持有读锁时，other goroutine 无法获取写锁，直到所有读锁都被释放。这保证了在 Broadcast 方法遍历 players 期间，Join 和 Leave 方法无法修改 players
*/

// 房间类
type Room struct {

	//一个map,key是玩家，value是bool,表示这个玩家是否在房间中（主要用key, 表示这个房间有多少玩家了）
	//由于map是并发不安全的，所以需要加锁
	players map[*Player]bool

	//读写互斥锁：读写锁允许多个 goroutine 同时进行读操作，但在写操作时会独占资源，防止其他 goroutine 进行读写操作，从而保证数据的一致性和完整性
	//在向 players 中添加或移除玩家（写操作）时，需要加写锁；在读取 players 中的玩家信息（读操作）时，需要加读锁。
	playersMutex sync.RWMutex

	//房间最大人数
	maxPlayers int
}

// 创建一个房间, 注意这里不是方法，是一个普通函数
func NewRoom() *Room {
	return &Room{
		players:    make(map[*Player]bool),
		maxPlayers: config.GlobalConfig.Game.MaxPlayers,
	}
}

// 加入房间
func (r *Room) Join(p *Player) {
	//这里加锁是必要的，因为遍历的player map的时候，有可能其他的goroutine在修改。map是并发不安全的。
	//加锁的目的是为了保证在遍历map的时候，其他的goroutine不会修改map。因为修改map的场景加了写锁。
	r.playersMutex.Lock() //	这里加锁是因为要操作map

	if len(r.players) >= r.maxPlayers {
		p.Send(protocol.NewErrorMessage("房间已满"))
		p.Conn.Close() //如果房间满了，就关闭连接
		return
	}
	r.players[p] = true

	r.playersMutex.Unlock() //这里需要释放锁，因为如果不释放会独占锁，在下面的Broadcast的时候会一直阻塞，造成死锁

	//广播通知其他人有新玩家加入
	r.Broadcast(protocol.NewSystemMessage("玩家加入房间, pid	= " + p.ID))
}

func (r *Room) Leave(p *Player) {
	r.playersMutex.Lock() //	这里加锁是因为要操作map
	delete(r.players, p)
	r.playersMutex.Unlock()

	//广播通知其他人有玩家离开
	r.Broadcast(protocol.NewSystemMessage("玩家离开房间, pid	= " + p.ID))
}

// 给房间内的所有玩家发送消息
func (r *Room) Broadcast(msg protocol.Message) {
	r.playersMutex.RLock() //这里加锁是因为要操作map
	defer r.playersMutex.RUnlock()
	for p, _ := range r.players {
		p.Send(msg)
	}
}

func (r *Room) HandlerMessage(sender *Player, msg protocol.Message) {
	switch msg.ID {
	case protocol.MsgHeartbeat:
		//心跳消息，不需要处理
		sender.UpdateLastHeartbeat()
	case protocol.MsgChat:
		//聊天消息，需要广播给其他玩家
		chatMsg := protocol.ChatMessage{
			From:    sender.ID,
			Content: msg.Data["content"].(string), //.(string)是强制类型转换的意思，这里是将msg.Data["content"]转换成string类型
		}
		r.Broadcast(protocol.NewMessage(protocol.MsgChat, chatMsg))
	default:

	}
}
