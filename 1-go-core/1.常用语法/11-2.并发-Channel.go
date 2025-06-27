package main

import "fmt"

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
}
