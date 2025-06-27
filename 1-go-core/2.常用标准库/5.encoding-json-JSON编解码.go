package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name  string `json:"name"` //tag-标签，json:"name" 表示json对象的key是name
	Age   int    `json:"age"`
	Email string `json:"email"`
}

/**
log.Fatal和panic的区别
相同点：终止程序并且打印日志
不同点：1.panic会打印堆栈，fatal不会 2.panic会执行defer, fatal不会
所以最好还是用panic
*/

func main() {
	p := new(Person)
	p.Name = "Jack"
	p.Age = 23
	p.Email = "cyou-inc.com"

	//将对象转换成json对象
	jsonData, err := json.Marshal(p) //Marshal-读音 妈手， 作用是把对象转换成json对象
	if err != nil {
		log.Fatal(err) //Fatal-读音-非特儿
	}
	fmt.Println(string(jsonData))

	//美化输出
	jsonPretty, err := json.MarshalIndent(p, "", " ") //prefix-前缀  indent-缩进（按什么来缩进）
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonPretty))

	//把json串解码成Person对象
	var p2 Person
	err2 := json.Unmarshal(jsonPretty, &p2) //第二个参数放的是指针
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(p2)

}
