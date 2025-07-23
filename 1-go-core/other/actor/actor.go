package main

import "sync"

/*
*
  - Actor模型
    1.可以接收消息
    2.可以处理消息
    3.可以订阅和处理事件
*/
type Actor struct {
	id      string
	inbox   chan Message //用于接收消息
	quit    chan struct{}
	handler func(actor *Actor, msg Message) //用于处理消息
	events  map[string][]func(interface{})  //简单的事件处理器,用于处理自己对自己的事件（应该有一个事件总线的，不应该在自己身上加事件相关处理单元） key=string value=[]func(interface{})是一个函数类型的切片,一个事件可以有多个处理函数
	mutex   sync.RWMutex                    //读写锁,用于防止并发读写
}

func NewActor(id string, handler func(actor *Actor, msg Message)) *Actor {
	return &Actor{
		id:      id,
		inbox:   make(chan Message),
		quit:    make(chan struct{}),
		handler: handler,
		events:  make(map[string][]func(interface{})),
	}
}

func (a *Actor) Start() {
	a.Publish("start", nil)
	go func() {
		defer a.Publish("end", nil)
		for {
			select {
			case msg := <-a.inbox:
				a.Publish("message", msg)
				a.handler(a, msg)
			case <-a.quit:
				return
			}
		}
	}() //匿名函数后面必须加()
}

func (a *Actor) Stop() {
	close(a.quit) //close方法是关闭channel的内置函数。close是一种广播机制，所有监听这个通道的goroutine收到通知。比发送数据（a.quit <- struct{}{}）更高效。
}

// 调用他的send可以给他发送消息， 他在start方法里会监听到你发给他的这个消息
func (a *Actor) Send(msg Message) {
	a.inbox <- msg
}

// 注册/订阅一个事件
func (a *Actor) Subscribe(event string, callback func(interface{})) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.events[event] = append(a.events[event], callback)
}

// 处理/发布一个事件
func (a *Actor) Publish(event string, data interface{}) {
	a.mutex.RLock()
	defer a.mutex.RUnlock()
	if listeners, ok := a.events[event]; ok {
		for _, fn := range listeners {
			go fn(data)
		}
	}
}

func (a *Actor) ID() string {
	return a.id
}
