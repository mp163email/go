package main

import "game-server/internal/network"

func main() {
	/**
	 	1.使用net.Listen监听tcp端口，等待客户端连接
		2.使用Accept接受连接
		3.使用bufferio.NewReader创建一个读取器，从连接中读取数据
		4.根据读取到的数据，调用不同的处理函数，处理不同的消息
		5.使用conn.Write向连接中写入数据
	*/
	network.StartServer()
}
