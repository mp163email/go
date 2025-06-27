package main

import "fmt"

func main() {
	d, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(d)
	}
}

func divide(a, b float64) (float64, error) { // error是一个基本类型
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero") //fmt.Errorf可以返回一个error类型的对象
	}
	return a / b, nil //空对象用nil表示
}
