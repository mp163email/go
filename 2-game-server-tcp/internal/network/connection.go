package network

import (
	"bufio"
	"fmt"
	"game-server/internal/game/handler"
	"net"
	"strings"
)

// 首字母大写，表示导出，包外的其他文件才可以调用
// 虽然这里首字母小写，但是当前包下的其他文件是可以调用的
func handlerConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn) //使用bufio.NewReader创建一个读取器，用于从连接中读取数据
	for {
		msg, err := reader.ReadString('\n') //阻塞读取数据，直到遇到换行符为止
		if err != nil {
			return
		}
		msg = strings.TrimSpace(msg)
		fmt.Println("Received message:", msg)
		if msg == "heartbean" {
			handler.HandlerHeartbeat(conn) //这里直接是用包.方法名调用, 没涉及到文件名
		} else if strings.HasPrefix(msg, "chat") {
			handler.HandlerChat(conn, msg)
		} else {
			conn.Write([]byte("Invalid command\n"))
		}
	}

}
