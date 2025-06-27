package main

import (
	"fmt"
	"math"
)

/**
接口需要结构体的方法来实现
*/

// 定义一个接口
type Shape interface {
	Area() float64
}

// 定义第一个结构体
type Circle struct {
	Radius float64
}

// 定义第二个结构体
type Square struct {
	Height float64
	Width  float64
}

// 给第一个结构体实现接口， 关键字是 func 后面的()代表所属
func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

// 给第二个结构体实现接口
func (s Square) Area() float64 {
	return s.Height * s.Width
}

func main() {
	var shap Shape = Circle{10} //把一个结构体的实例赋值给这个接口
	area := shap.Area()         //接口调用自己的方法
	fmt.Println(area)

	//断言-判断接口类型是否为指定的类型
	circle, ok := shap.(Circle)
	fmt.Println(circle, ok)

	//判断这个接口的实现类是哪个
	if circle, ok := shap.(Circle); ok {
		fmt.Println(circle.Radius)
	} else {
		fmt.Println("<UNK>")
	}

	//使用switch判断是哪个实现类
	square := Square{Height: 10, Width: 10}
	printShape(square)
	printShape(circle)

}

// 使用switch判断类型并使用里面的值
func printShape(s Shape) {
	switch s := s.(type) {
	case Circle: //这里的s代表类型也代表值， 和case用的时候判断的是s是什么类型, 和s.Area使用的时候，他代表的是具体的实例值
		fmt.Println("Circle, area=", s.Area())
	case Square:
		fmt.Println("Square, area=", s.Area())
	}
}
