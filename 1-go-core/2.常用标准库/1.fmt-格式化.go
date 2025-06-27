package main //main包是一个特殊的包，只有main包才能有main方法

import "fmt"

/*
*
方法名首字母是大写的
*/
func main() {
	fmt.Print("Go runs on ")
	fmt.Println("hello world")
	name := "miao"
	age := 25
	fmt.Printf("name=%s, age=%d\n", name, age) // Printf格式化

	//字符串格式化, Sprintf方法
	s := fmt.Sprintf("name is %s, age is %d\n", name, age)
	fmt.Println(s)

	//输入 Scanln方法,传入一个变量的指针地址
	var input string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&input)
	fmt.Println("Hello " + input)
}
