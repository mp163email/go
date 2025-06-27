package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(math.Sqrt(16))   //开平方
	fmt.Println(math.Pow(2, 3))  //2的3次方
	fmt.Println(math.Max(1, 2))  //求最大值
	fmt.Println(math.Min(1, 2))  //求最小值
	fmt.Println(math.Abs(-3.14)) //求绝对值
	fmt.Println(math.Round(3.6)) //四舍五入
	fmt.Println(math.Ceil(3.2))  //向上取整
	fmt.Println(math.Floor(3.2)) //向下取整

	fmt.Println(math.Sin(math.Pi / 2)) //三角函数

	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(100)) //0-99之间的随机数
	fmt.Println(rand.Float64()) //0.0-1.0之间的随机浮点数

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})
	fmt.Println(numbers)

}
