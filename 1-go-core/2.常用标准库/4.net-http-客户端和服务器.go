package main

import (
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
}
