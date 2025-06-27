package main

import "fmt"

func main() {
	x := add(2, 3) // :=是变量的声明+赋值   单独赋值的话,不能加:
	fmt.Println(x)

	x = multiply(2, 3)
	fmt.Println(x)

	var a, b, c string = swap("a", "b") //用多个变量来接,不需要用数组
	fmt.Println(a, b, c)

	f, g := split(10)
	fmt.Println(f, g)

	sum := sum(1, 2, 3)
	fmt.Println(sum)
}

// 带参数和返回值的函数, 数据类型位置放到后面
func add(x int, y int) int {
	return x + y
}

// 函数参数的简写(如果两个参数的类型一致)
func multiply(x, y int) int {
	return x * y
}

// 很有特色，非常实用的语法：真没想到, go的函数可以返回多个值
func swap(x, y string) (string, string, string) {
	return y, x, x
}

// 可以在函数上定义好返回哪几个变量名
func split(sum int) (x, y int) {
	x = sum * 4 / 9 //这里的x, y直接赋值
	y = sum - x
	return //函数上面已经定义了,return的时候就能够省掉
}

// 可变参数的函数
func sum(num ...int) int {
	total := 0
	for _, i := range num {
		total += i
	}
	return total
}
