package main

import "fmt"

func main() {
	x := 2
	p := &x // &代表取地址， 地址是用来赋值给指针的
	y := *p // *代表取值，也叫解引用
	fmt.Println(y)
	*p = 20 //可以通过*p来改变指向地址的值
	fmt.Println(x)

	//利用指针实现方法的引用传递（默认是值传递）
	var pointer *int
	var value = 6
	pointer = &value
	fmt.Println(pointer, *pointer)

	refer(&value)
	fmt.Println(value)

}

// 把变量的地址（指针）传过来
func refer(x *int) {
	*x = 100 //使用*修改指针指向的值
}
