package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

func main() {
	//读取一个字符串到字节数组
	reader := strings.NewReader("Hello, Reader!")
	buf := make([]byte, 8) //定义一个8字节的缓存,每次读8个字节做为缓冲
	var all []byte
	for {
		read, err := reader.Read(buf)
		//append 追加数据
		all = append(all, buf[:read]...) //...是展开操作符, 表示吧buf[:read]中每个字节展开，并逐个追加到all里
		fmt.Printf("n = %v, err = %v, buf=%v\n", read, err, buf)
		if err == io.EOF {
			break
		}
	}
	fmt.Println(string(all))

	//读取一个文件  方法被弃用，建议用os标准库来操作[os.ReadFile]
	file, err := ioutil.ReadFile("设计哲学")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(file))

	//写入一个文件
	message := []byte("hello world")                     //将一个字符串转换成字节数组, 这里可以直接放一个字符串
	err1 := ioutil.WriteFile("test1.txt", message, 0666) //0666是文件的权限设置 （可读可写）
	if err1 != nil {
		panic(err1)
	}
}
