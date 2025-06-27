package main

import "fmt"

/**
函数和方法的区别：
	1.所属： 方法是绑定到某个类型上的函数，而普通函数是独立的与类型无关
	2.调用： 方法需要结构体对象来调用, 函数随便用
	3.书写： 方法有所属，在func 后面的（）里, 函数没有
函数是孤立的工具-类似java中的静态方法，方法是结构体的行为（类似java中类的方法）
*/

// 长方形结构体（类）
type Rectangle struct {
	Width, Height float64
}

// 方法-求面积 这里直接使用的属性的值
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

// 方法-缩放长和宽 这里改变属性的值,直接用的指针
func (r *Rectangle) Scale(factor float64) {
	r.Height *= factor
	r.Width *= factor
}

func main() {
	//p := Rectangle{3, 4}//构造方法的方式
	p := new(Rectangle) //new的方式
	p.Height = 3
	p.Width = 4
	area := p.Area() //调用的时候,需要用类型来调用，计算的时候也是使用的类型里面的值
	fmt.Println(area)

	p.Scale(2) //修改属性的值
	fmt.Println(p.Area())
}
