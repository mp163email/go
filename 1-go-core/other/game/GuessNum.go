package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

/*
* 猜数字
GOROOT=C:\Program Files\Go #gosetup
GOPATH=C:\Users\miaopeng\go #gosetup
"C:\Program Files\Go\bin\go.exe" build -o C:\Users\miaopeng\AppData\Local\JetBrains\GoLand2025.1\tmp\GoLand\___go_build_1_go_core_other_game.exe 1-go-core/other/game #gosetup
C:\Users\miaopeng\AppData\Local\JetBrains\GoLand2025.1\tmp\GoLand\___go_build_1_go_core_other_game.exe #gosetup
欢迎来到猜数字游戏！
我已经想好了一个1到100之间的数字，你可以开始猜了。
请输入你的猜测：50
太小了，再试一次
请输入你的猜测：75
太小了，再试一次
请输入你的猜测：90
太大了，再试一次
请输入你的猜测：80
太小了，再试一次
请输入你的猜测：85
太小了，再试一次
请输入你的猜测：87
恭喜你，你猜对了！你用了6次机会。
*/
func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano())) // new 一个Random
	target := r.Intn(100) + 1
	fmt.Println("欢迎来到猜数字游戏！")
	fmt.Println("我已经想好了一个1到100之间的数字，你可以开始猜了。")

	reader := bufio.NewReader(os.Stdin) //new一个输入流的reader
	attempts := 0

	for {
		attempts++
		fmt.Print("请输入你的猜测：")
		input, err := reader.ReadString('\n') //调用ReadString方法, 读取到\n为止
		if err != nil {
			fmt.Println("读取输入出错")
			continue
		}
		input = input[:len(input)-1]
		guess, err := strconv.Atoi(input) //字符串转整数
		if err != nil {
			fmt.Println("请输入有效的数字")
			continue
		}
		if guess < 1 || guess > 100 {
			fmt.Println("请输入1-100之间的数字")
			continue
		}
		if guess < target {
			fmt.Println("太小了，再试一次")
		} else if guess > target {
			fmt.Println("太大了，再试一次")
		} else {
			fmt.Printf("恭喜你，你猜对了！你用了%d次机会。\n", attempts)
			break
		}
	}
}
