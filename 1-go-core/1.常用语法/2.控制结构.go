package main

import (
	"fmt"
	"runtime"
)

func main() {

	//if-else if - else语句
	x := 3
	if x > 1 {
		fmt.Println("x is greater than 1")
	} else if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is equal to 5")
	}
	//if 加默认函数赋值
	if y := calcIf(); y < 8 {
		fmt.Println("Y is less than 8")
	}

	//switch
	switch os := runtime.GOOS; os { //runtime是go标准库的一个包,表示程序运行的操作系统
	case "windows":
		fmt.Println("os is windows")
	case "linux":
		fmt.Println("os is linux")
	default:
		fmt.Println("os is 3.other, os = ", os)
	}

	// 循环
	// for
	for i := 0; i < 10; i++ { //for后面没有()
		fmt.Println("当前value=", i)
	}
	// while的替代
	sum := 1
	for sum < 10 {
		sum += 2
		fmt.Println("curr value=", sum)
	}
	// 无限循环 while (true)的替代
	num1 := 0
	for {
		num1 += 1
		if num1 > 10 {
			break
		}
		fmt.Println("<UNK> num1=", num1)
	}
	//遍历数组
	nums := []int{1, 2, 3, 4, 5}
	for i, num := range nums { // 1.range是遍历集合的一个关键字, 2.这里用的是,号, 3.range前面有 := 4.for后没有()
		fmt.Printf("%d\t%d\n", i, num)
	}
	str := "Go"
	for i, char := range str { // range是遍历集合的一个关键字
		fmt.Printf("%d\t%c\n", i, char)
	}
}

func calcIf() int {
	return 5
}
