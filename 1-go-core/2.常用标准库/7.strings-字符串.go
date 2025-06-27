package main

import (
	"fmt"
	"strings"
)

func main() {
	//包含
	fmt.Println(strings.Contains("hello world", "hello"))

	//统计某个字符/字符串 再源字符串出现的个数
	fmt.Println(strings.Count("hello world", "o"))

	// 分割
	fmt.Printf("%q\n", strings.Split("a,b,c", ",")) // %q的意思是把切片里的每个数据用 ""括起来

	//连接
	fmt.Println(strings.Join([]string{"a", "b", "c"}, "-"))

	//替换 最后一个参数是替换几个
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 1))

	//大小写转换
	fmt.Println(strings.ToUpper("hello world"))
	fmt.Println(strings.ToLower("HELLO WORLD"))

	//修剪
	fmt.Println(strings.TrimSpace("   hello   world")) //去空格
	fmt.Println(strings.Trim("!!Hello!!!", "!"))

	//是否是什么开头（前缀）, 是否是什么结尾（后缀）
	fmt.Println(strings.HasPrefix("Hello World", "H"))
	fmt.Println(strings.HasSuffix("hello", "l"))
}
