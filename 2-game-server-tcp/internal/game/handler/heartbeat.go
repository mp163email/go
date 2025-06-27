package handler

import (
	"net"
	"time"
)

// 处理心跳消息的函数
// 心跳消息是客户端发送给服务器的一个特殊消息，用于保持连接的活跃状态
func HandlerHeartbeat(conn net.Conn) {
	conn.SetDeadline(time.Now().Add(10 * time.Second)) //设置读写超时时间 10秒
	conn.Write([]byte("pong \n"))
}
