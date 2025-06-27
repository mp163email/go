package network

import (
	"fmt"
	"game-server/internal/config"
	"net"
)

func StartServer() {
	// 启动服务器
	ln, err := net.Listen("tcp", config.ServerPort) //使用net.Listen函数创建一个TCP监听器，监听指定的端口
	if err != nil {
		panic(err)
	}
	fmt.Println("server listening on", config.ServerPort)

	for {
		conn, err := ln.Accept() //阻塞等待客户端连接
		if err != nil {
			continue
		}
		go handlerConnection(conn) //开启一个goroutine处理连接
	}
}
