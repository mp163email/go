package main

import "fmt"

func main() {
	//使用make关键字创建一个map,只定义
	m := make(map[string]int) //注意这里只有key用了[]
	m["name"] = 1
	m["age"] = 18
	fmt.Println(m, m["name"])

	//定义+赋值 字面量创建一个map, 需要初始化
	mp := map[string]string{
		"name": "lucy",
		"age":  "18",
	}
	fmt.Println(mp)

	//获取map中某个key的值
	v, exits := m["name"]
	fmt.Println(v, exits)

	//获取map中某个key的值, 其实返回了2个值，第一个是这个key的值，第二个是是否存在这个key
	val, ok := m["aaaaa"]
	if ok {
		fmt.Println(val)
	} else {
		fmt.Println("not found")
	}

	//使用delete关键字删除map中的某个key
	delete(m, "name")
	fmt.Println(m)

	//使用range遍历map
	for k, v := range m {
		fmt.Printf("k=%s, v=%d", k, v)
	}
}
