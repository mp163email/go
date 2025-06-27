package main

import "fmt"

func main() {
	var i int = 3
	var f float64 = float64(i)
	var u uint = uint(f) //无符号整数 去掉小数部分,保留整数部分（向下取整）
	fmt.Println(i, f, u)
}
