package main

import (
	"fmt"
	"time"
)

/**
典型的Coroutine+Channel的运用
案例：将一个数组分层2部分， 各部分交给各自的协成并发执行（两个协程同时计算，提升执行效率）， 各自执行的结果写入到chanel里， 从channel里取出计算的值
并行计算数组的和
channel： 一个用于在不同 goroutine（协程）之间传递数据的“通信通道” 简写chan是缩写
*/

func sum1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum //把结果发送到channel
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int)      //创建一个能接收int类型的channel
	go sum1(s[:len(s)/2], c) //把数组分成左右2部分 :左边是开始下标, 右边是结束下标,左侧没有值=0, 右侧没有值=len
	go sum1(s[len(s)/2:], c)
	x, y := <-c, <-c       //从channel中取出数据来
	fmt.Println(x, y, x+y) //打印值

	cc := make(chan int)
	go func() {
		fmt.Println("cc 发送方准备发送数据: 1")
		cc <- 1                               //走到这行, 会阻塞, 等待接收方接收
		fmt.Println("cc 发送方完成发送,这行会在接收后才会执行") // 这行会在接收后才会执行
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("cc 接收方接收数据:", <-cc)

	/**
	带缓冲Channel就像一个有容量的队列
	发送数据时只要队列不满就不会阻塞
	接收数据时只要队列不空就不会阻塞
	缓冲大小决定了可以"提前"发送多少数据而不用等待接收
	合理设置缓冲大小可以提高程序性能，避免频繁阻塞
	*/
	//带缓冲的channel, 不用阻塞
	ccc := make(chan int, 2)
	go func() {
		defer func() {
			if err := recover(); err != nil { //捕获和处理panic异常
				fmt.Printf("cc 协程发生panic: %v\n", err)
			}
		}()
		fmt.Println("ccc 给带缓冲的channel发送一个数")
		ccc <- 1
		ccc <- 1
		cc <- 1 //这里goroutine会泄漏，因为它永远无法完成（接收方只有1个）
		fmt.Println("ccc 发送完数据,但是不用阻塞")
	}()

	time.Sleep(2 * time.Second)
	received := <-ccc
	fmt.Println("ccc 从带缓冲的channel接收一个数", received)
}
