package main

import (
	"fmt"
	"time"
)

func main() {

	//当前时间
	now := time.Now()
	fmt.Println(now) //2025-06-13 15:45:20.4989596 +0800 CST m=+0.000000001

	//格式化时间
	fmt.Println(now.Format("2006-01-02 15:04:05")) //Go团队设计的一个格式模版，不能换成别的。用实例表达规则

	//字符串转时间 time.Parse方法
	parse, err := time.Parse("2006-01-02", "2025-06-13") //parse是一个Time类型的数据
	if err != nil {
		panic(err)
	}
	fmt.Println(parse)

	//时间运算
	tomorrow := now.Add(24 * time.Hour)
	fmt.Println(tomorrow)

	//求时间间隔 Since-单位-秒
	start := time.Now()
	time.Sleep(1 * time.Second)
	elapsed := time.Since(start)
	fmt.Println(elapsed)

	//定时器 time.NewTimer, 和time.Sleep相比， 前者可以被打断，功能更多
	timer := time.NewTimer(2 * time.Second)
	<-timer.C //意思是“从 timer.C 通道中接收一个数据，但不保存它”  代码会阻塞（暂停执行），直到通道里有数据可接收
	fmt.Println("Timer fired")

	//tick
	ticker := time.NewTicker(2 * time.Second) //每隔2秒会向ticker.C通道（channel)发数据
	go func() {
		for t := range ticker.C {
			fmt.Println("Ticker fired", t)
		}
	}()
	time.Sleep(10 * time.Second)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
