package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//带值的上下文管理 //类似一个map，key-value
	valueCtx := context.WithValue(context.Background(), "miao", "peng")
	fmt.Println(valueCtx.Value("miao"))

	//取消上下文管理 WithCancel
	cancelCtx, cancel := context.WithCancel(context.Background())
	go worker(cancelCtx, 1)
	go worker(cancelCtx, 2)
	time.Sleep(time.Second * 3)
	cancel() //cancel方法
	time.Sleep(time.Second)

	//超时上下文管理 WithTimeout
	timeoutCtx, cancelTimeOut := context.WithTimeout(context.Background(), time.Second*2)
	defer cancelTimeOut()
	go worker(timeoutCtx, 1)
	time.Sleep(time.Second * 3) //睡3秒，但是只执行了2秒，因为2秒时超时了
}

func worker(ctx context.Context, id int) {
	//采用while true的方式来实现一直工作，直到被取消. select是用来监听多个通道的,如果有一个通道有数据，就会执行对应的case
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d exit\n", id)
			return
		default:
			fmt.Printf("worker %d is working\n", id)
		}
	}
}
