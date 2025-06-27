package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//等待组
	var wg sync.WaitGroup //创建一个等待组  等待组

	for i := 0; i < 3; i++ {
		wg.Add(1) //增加等待组中的goroutine数量
		go func(i int) {
			defer wg.Done() //goroutine完成后减少等待组中的goroutine数量
			time.Sleep(time.Duration(i) * time.Second)
			fmt.Println(i)
		}(i)
	}

	wg.Wait() //等待所有goroutine完成
	fmt.Println("All goroutines finished")

	//互斥锁
	var mu sync.Mutex //创建一个互斥锁
	var counter int
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock() //加锁 加锁时其他的goroutine就不能拿到这个锁,会阻塞等待，直到锁被释放。确保只有一个goroutine在访问这个变量
			counter++
			mu.Unlock() //解锁
		}()
	}
	wg.Wait()
	fmt.Println("All goroutines finished", counter)

	//只执行一次
	var once sync.Once
	setup := func() {
		fmt.Println("setup")
	}
	for i := 0; i < 10; i++ {
		once.Do(setup) //虽然for循环10次，但是只会执行一次setup函数
	}
}
