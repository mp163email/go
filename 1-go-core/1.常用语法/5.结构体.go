package main

import "fmt"

// 结构体 type关键字 + struct关键字
// 相当于Java中的类
// 结构体首字母大写-可以被其他包访问,需要import  首字母小写-只能在同一个包中使用
type Person struct {
	Name string
	Age  int
}

func main() {

	//使用隐式构造函数构建，默认和结构体定义的字段顺序一致
	p := Person{"Alice", 25} //没有new关键字  用的是{}
	fmt.Println(p.Name, p.Age)

	//使用字段名赋值  注意这里用的是:号
	p2 := Person{Age: 18, Name: "Lucy"}
	fmt.Println(p2.Age, p2.Name)

	//结构体指针,可以自动解引用,不用加*
	pp := &p2
	fmt.Println(pp.Age, pp.Name)

	//匿名结构体
	nm := struct { //没有type关键字, 只有struct关键字
		X, Y int
	}{10, 20} //默认赋值是必须的
	fmt.Println(nm, nm.X, nm.Y)
}
