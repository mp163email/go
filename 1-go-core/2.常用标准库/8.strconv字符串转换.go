package main

import (
	"fmt"
	"strconv"
)

func main() {
	//字符串转换成整数
	atoi, err := strconv.Atoi("45")
	if err != nil {
		panic(err)
	}
	fmt.Println(atoi)

	//整数转字符串
	itoa := strconv.Itoa(45)
	fmt.Println(itoa)

	//字符串转bool
	parseBool, err := strconv.ParseBool("true")
	if err != nil {
		panic(err)
	}
	fmt.Println(parseBool)

	//字符串转float
	mfloat, err := strconv.ParseFloat("3.1415", 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(mfloat)

	//格式化浮点数
	mmfloat := strconv.FormatFloat(3.1415926, 'f', 3, 64)
	fmt.Println(mmfloat)

}
