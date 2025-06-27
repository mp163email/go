package main

import (
	"fmt"
	"time"
)

/**
Select: 用于监听多个 channel 操作的一种机制
Select: 从多个 channel 操作中选择一个可以执行

 特性	            说明
多路监听	    同时监听多个 channel
非阻塞组合	可以配合 default 分支实现非阻塞操作
随机选择	    如果多个 channel 同时就绪，select 会随机选一个

用的是select-case结构
*/

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- 1
	}()
	go func() {
		time.Sleep(5 * time.Second)
		ch2 <- 2
	}()
	for i := 0; i < 2; i++ {
		select { //select 会阻塞 等待ch1和ch2返回值
		case v1 := <-ch1:
			fmt.Println("v1 = ", v1)
		case v2 := <-ch2:
			fmt.Println("v2 = ", v2)
			//default: //不阻塞,直接往下走
			//	fmt.Println("default")
		}
	}
}
