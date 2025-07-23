package main

import (
	"fmt"
	"time"
)

func main() {
	textHandler := func(a *Actor, msg Message) {
		switch msg := msg.(type) {
		case *TextMessage:
			fmt.Printf("[%s] Received:%s\n", a.ID(), msg.Content)
		default:
			fmt.Printf("[%s] Received:%s\n", a.ID(), msg.Type())
		}
	}

	a1 := NewActor("monster-1", textHandler)
	a2 := NewActor("player-1", textHandler)

	//给a1注册事件监听
	a1.Subscribe("start", func(data interface{}) {
		fmt.Println("event-start, id=", a1.ID())
	})
	a2.Subscribe("message", func(data interface{}) {
		if msg, ok := data.(*TextMessage); ok {
			fmt.Printf("[%s] event-message-received:%s\n", a2.ID(), msg.Content)
		}
	})

	a1.Start()
	a2.Start()

	a1.Send(&TextMessage{Content: "Monster is angry!"}) //main方法里,调用a1的send方法,向a1的inbox通道发送一个消息
	a2.Send(&TextMessage{Content: "Player attacks!"})   //main方法里,调用a2的send方法,向a2的inbox通道发送一个消息

	time.Sleep(1 * time.Second)

	a1.Stop()
	a2.Stop()
}
