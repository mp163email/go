package main

import (
	"flag"
	"fmt"
	"http-server/internal/app"
	"http-server/internal/hello"
	"http-server/internal/router"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
*
提供如下访问接口: http://localhost:8081/readyz, http://localhost:8081/api/v1/hello等
[GIN-debug] GET    /healthz                  --> http-server/internal/router.New.func1 (5 handlers)
[GIN-debug] GET    /readyz                   --> http-server/internal/router.New.func2 (5 handlers)
[GIN-debug] GET    /version                  --> http-server/internal/router.New.func3 (5 handlers)
[GIN-debug] GET    /api/v1/hello/ping        --> http-server/internal/hello.(*Handler).ping-fm (5 handlers)
[GIN-debug] GET    /api/v1/hello             --> http-server/internal/hello.(*Handler).greet-fm (5 handlers)
*/
func main() {

	//命令行参数
	addr := flag.String("addr", ":8081", "http server listen address")
	flag.Parse()

	//构建handler和service
	helloSrvice := hello.NewHelloService()
	helloHandler := hello.NewHandler(helloSrvice)

	//构建一个gin引擎, 它实现了http.Handler接口,并构建了路由, 可以直接用在http.ListenAndServe中使用
	r := router.New(router.Options{
		HelloHandler: helloHandler,
	})

	//构建一个http服务器, 并将gin引擎作为它的处理器
	srv := app.NewHttpServer(r, *addr)

	//异步启动http服务器
	go func() {
		if err := app.Start(srv); err != nil {
			fmt.Println("start err:", err)
		}
	}()
	fmt.Printf("http server start at %s\n", *addr)

	//等待信号量, ctrl+c / kill
	quit := make(chan os.Signal, 1)                      // 创建一个容量为1的信号通道，用来接收系统发来的中断信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 注册监听两种信号：SIGINT（Ctrl+C）和 SIGTERM（kill命令默认信号）
	<-quit                                               // 阻塞等待信号到来（程序会在这里卡住，直到收到信号）
	_ = app.Stop(srv, 10*time.Second)                    // 收到信号后开始执行优雅关闭流程，最多等待10秒
	fmt.Println("app stop")
}
