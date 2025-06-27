package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//从标准输入读取数据
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter text: ")
	scanner.Scan()
	text := scanner.Text()
	fmt.Print("Your enter text: ", text)

	//从字符串读取数据
	fmt.Println()
	reader := strings.NewReader("Hello, World!")
	newScanner := bufio.NewScanner(reader)
	for newScanner.Scan() {
		fmt.Println(newScanner.Text())
	}

	//缓冲区
	create, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer create.Close()
	writer := bufio.NewWriter(create)
	writer.WriteString("Hello, World!\n") //写入缓冲区
	writer.Flush()
}
