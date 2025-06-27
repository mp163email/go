// main包必须包含一个main()函数做为入口点
package main

//导入包
import "fmt"

// 行尾不需要加;
// go的变量定义,把类型放到变量名的后面
var age int = 18          //var 定义的变量可以定义在函数以外, var必不可少, 不用:=中的: 就得用var
var name string = "hello" // string 中的s是小写

// 定义一个新的类型MsgID，int 是他的基础类型
type MsgId int

func main() {

	var id MsgId = 111 //给类型直接赋值
	fmt.Println(id)

	var num int = 222
	id = MsgId(num) //需要类型转换, 假如直接id = num 会报错, 因为num是int类型, id是MsgId类型, 不能直接赋值
	fmt.Println(id)

	var isOpen = true                        //使用了var, 就不能再使用:=类型推断
	var str1, str2 string = "hello", "world" //多变量定义
	fmt.Println(name, age, isOpen, str1, str2)

	//类型推断
	count := 42              // :=只能定义在函数内, 不需要var和显示类型. :=中的:表示声明变量 =表示给变量赋值, 如果给变量二次赋值，不能再用:=因为:已经声明过了，只需要用=
	a, b := "hello", "world" // 多变量定义
	fmt.Println(count, a, b)

	//其他常用数据类型
	var isActive bool = true
	var price float32 = 19.9
	fmt.Println(isActive, price)

	//单个常量
	const Pi float64 = 3.1415926
	//常量组
	const (
		OK       = 200
		NotFound = 404
	)
	fmt.Println(Pi, OK, NotFound)

	//数组
	var xxx = [5]int{1, 2, 3, 4, 5}
	yyy := []int{6, 7, 8, 9, 10}
	fmt.Printf("第一个数=%d, 第2个数=%d, 第3个长度=%d", xxx[2], yyy[0], len(yyy)) //输出带格式的用Printf方法

	//空接口 - 可以接收任意类型的值
	var strs string = "hello"
	var any interface{} //任何的类型都实现了空接口，所以我们可以把任何一个变量赋给空接口
	any = strs
	fmt.Println("\n", any)
}
