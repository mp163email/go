package main

import (
	"fmt"
	"os"
)

func main() {
	//Getenv方法获取环境变量
	getenv := os.Getenv(`PATH`)
	fmt.Printf("getenv = %s\n", getenv)

	//Setenv方法设置环境变量
	os.Setenv("MY_VAR", "some value")
	fmt.Printf("MY_VAR = %s\n", os.Getenv("MY_VAR"))

	//os.Args获取命令行参数
	args := os.Args
	fmt.Println("命令行参数=", args) //自身路径：[C:\Users\miaopeng\AppData\Local\JetBrains\GoLand2025.1\tmp\GoLand\___go_build_2_os__go.exe]  启动时传入的参数

	//创建文件 os.Create
	file, err := os.Create("test.txt")
	if err != nil { //比如权限不足导致文件创建失败
		panic(err) //panic代表崩溃终止程序并打印错误 类似Java中的throw new RunTimeException(err)
	}
	file.WriteString("Hello World")
	file.Close()

	//读取文件 os.ReadFile
	data, err := os.ReadFile("test.txt") //看源码发现data的类型是[]byte
	if err != nil {
		panic(err)
	}
	fmt.Println("File Content=", string(data)) //类型转换-string(data)-将字节数组转换成string

	//删除文件
	erro := os.Remove("test.txt")
	if erro != nil {
		fmt.Println(erro)
	}
}
