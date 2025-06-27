package pojo

import "github.com/lonng/nano"

// 房间对象， 就包含一个group对象
type Room struct {
	Group *nano.Group //Group是nano的核心组件(它自己封装的)，用于管理会话（sessions-map结构）和广播消息（遍历sessions-s.Push）
}

type AllMembers struct {
	Members []int64 `json:"members"`
}
