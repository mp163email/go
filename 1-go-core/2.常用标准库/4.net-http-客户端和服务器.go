package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// 一个处理函数
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello %s!", r.URL.Path[1:])
}

func main() {
	//绑定url和处理函数
	http.HandleFunc("/hello/", helloHandler)

	// 启动 HTTP 服务，并捕获错误
	go func() {
		err := http.ListenAndServe(":8081", nil)
		if err != nil {
			panic(err)
		}
	}()

	time.Sleep(5 * time.Second)

	//模拟一个客户端请求
	resp, err := http.Get("http://127.0.0.1:8081/hello")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("All Content=", string(all))

	///////////////////////使用http.Get方法调用一个链接/////////////////
	fmt.Printf("Http请求实例")
	resp, err1 := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err1 != nil {
		fmt.Println("Http请求失败", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("相应状态码", resp.StatusCode)

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println("JSON解析错误", err)
		return
	}
	fmt.Println("帖子标题:", result["title"])

}
