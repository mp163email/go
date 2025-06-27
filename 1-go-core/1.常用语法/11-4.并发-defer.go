package main

import (
	"fmt"
	"time"
)

/*
*
┌─────────────────────────────────────┐
│              defer 用法图解        │
├─────────────────────────────────────┤
│ ❶ 定义语法：                        │
│     defer 函数/语句                 │
│                                     │
│ ❷ 执行时机：                        │
│     当前函数返回前执行              │
│                                     │
│ ❸ 执行顺序：                        │
│     后进先出（LIFO）                │
│                                     │
│ ❹ 参数求值：                        │
│     在 defer 定义时立即求值        │
├─────────────────────────────────────┤
│ ❺ 常见用途：                        │
│   - 关闭文件：     defer f.Close()  │
│   - 解锁：        defer mu.Unlock()│
│   - 捕获异常：     defer recover()  │
│   - 统计耗时：     defer log耗时    │
└─────────────────────────────────────┘
函数退出前执行这个预定义的函数(延迟执行预定义的函数)
可以统计耗时，关闭文件等
*/
func hello() {
	defer fmt.Println("world") //延迟到hello方法退出前才执行
	fmt.Println("hello")
}

func main() {
	start := time.Now()

	//统计main函数的总耗时
	defer func() { //******最先注册 defer，因此最后执行， 按代码顺序决定注册顺序*****
		fmt.Println("cost :", start, time.Since(start))
	}()

	hello()

	for i := 0; i < 10; i++ {
		defer fmt.Println(i) //后进先出， 顺序是倒着来的
	}

}
