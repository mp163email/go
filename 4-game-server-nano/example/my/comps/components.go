package comps

import (
	"fmt"
	"game-server-nano/example/my/cst"
	"game-server-nano/example/my/msgs"
	"game-server-nano/example/my/pojo"
	"github.com/lonng/nano"
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/pipeline"
	"github.com/lonng/nano/scheduler"
	"github.com/lonng/nano/session"
	"time"
)

// 房间管理器组件，有一个存储所有房间的map
type RoomManager struct {
	component.Base // 内嵌Base组件，提供组件生命周期管理
	timer          *scheduler.Timer
	rooms          map[int]*pojo.Room // map结构-存储所有房间
}

func NewRoomManager() *RoomManager {
	return &RoomManager{
		rooms: make(map[int]*pojo.Room),
	}
}

// 玩家加入房间 s *session.Session代表新加入的玩家
// 家里不论所有人，加入的都是同一个房间-默认房间roomId=1
func (mgr *RoomManager) Join(s *session.Session, msg []byte) error {
	//找默认房间roomId=1的是否存在，不存在就创建一个
	room, found := mgr.rooms[cst.TestRoomId]
	if !found {
		room = &pojo.Room{
			Group: nano.NewGroup(fmt.Sprintf("room-%d", cst.TestRoomId)),
		}
		mgr.rooms[cst.TestRoomId] = room
	}

	//给session uid赋值
	s.Bind(s.UID())            //有点多此一举，因为session的uid就是s.UID()
	s.Set(cst.RoomIdKey, room) //给session 的data赋值

	// 向新加入的当前用户推送当前房间内所有成员的信息， 客户端收到onMemebers消息后，会调用onMemebers回调函数, route这里指的是客户端的route
	s.Push("onMembers", &pojo.AllMembers{
		Members: room.Group.Members(),
	})

	//给房间内所有人推送新用户加入的消息
	room.Group.Broadcast("onNewUser", &msgs.NewUserMessage{
		Content: fmt.Sprintf("New User：%d", s.ID()),
	})

	// 加入房间
	room.Group.Add(s)

	//s.Response常用来对请求就行结果反馈， s.Push主动想客户端推送消息
	return s.Response(&msgs.JoinResponseMessage{Code: 0, Result: "success"})
}

// 有人发消息了，就把这条消息广播给房间内的所有人
func (mgr *RoomManager) Message(s *session.Session, message *msgs.UserMessage) error {
	if !s.HasKey(cst.RoomIdKey) {
		return fmt.Errorf("not found RoomIdKey") //fmt.Errorf可以用来创建一个error对象
	}
	room := s.Value(cst.RoomIdKey).(*pojo.Room)
	return room.Group.Broadcast("onMessage", message)
}

func (mgr *RoomManager) AfterInit() {
	session.Lifetime.OnClosed(func(s *session.Session) {
		if !s.HasKey(cst.RoomIdKey) {
			return
		}
		room := s.Value(cst.RoomIdKey).(*pojo.Room)
		room.Group.Leave(s) //连接断开的时候，把连接session从group 的sessions中删除
	})

	// scheduler是nano的定时器组件
	//每分钟给打印各个房间内有多少人（虽然这里只有一个房间）
	scheduler.NewTimer(time.Minute, func() {
		for roomId, room := range mgr.rooms {
			fmt.Printf("RoomId: %d, MemberCount: %d\n", roomId, room.Group.Count())
		}
	})
}

// 统计组件
type Status struct {
	component.Base
	timer         *scheduler.Timer
	outboundBytes int
	inboundBytes  int
}

func (stats *Status) Outbound(s *session.Session, msg *pipeline.Message) error {
	stats.outboundBytes += len(msg.Data)
	return nil
}

func (stats *Status) Inbound(s *session.Session, msg *pipeline.Message) error {
	stats.inboundBytes += len(msg.Data)
	return nil
}

func (stats *Status) AfterInit() {
	stats.timer = scheduler.NewTimer(time.Minute, func() {
		println("OutboundBytes:", stats.outboundBytes) //开发调试使用println, 生产环境使用log或者fmt
		println("InboundBytes:", stats.inboundBytes)
	})
}
