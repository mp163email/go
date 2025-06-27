package main

import (
	"log"
	"os"
)

func main() {
	//基本使用 时间是默认带的 2025/06/20 15:11:18 hello world
	log.Println("hello world")

	//带前缀的日志
	log.SetPrefix("MyApp:")
	log.Println("hello world with prefix")

	//带标记的日志
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("hello world with flags")

	//创建一个文件
	file, err := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	//将这个文件作为日志的输出
	logger := log.New(file, "Prefix:", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("Hello World")
}
