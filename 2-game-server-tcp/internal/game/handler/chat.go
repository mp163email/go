package handler

import (
	"fmt"
	"net"
)

// 处理聊天消息的函数
func HandlerChat(conn net.Conn, message string) {
	fmt.Printf("[Chat]%s \n", message)
	conn.Write([]byte("[Server] Chat received: " + message + "\n"))
}
