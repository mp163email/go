```
1.使用http.Server启动http服务
2.但是http组装handler的时候，使用gin框架(gin.Engine实现了handler的接口),来实现了访问路由
3.gin相关语法, gin的3个重要组件 1.Context-上下文 2.RouterGroup-路由 3.middleware中间件

 c *gin.Context 上下文，用来处理请求和响应
    c.Request 原始的http.Request
    c.Response 原始的http.ResponseWriter
    c.Next() 调用下一个中间件
    c.Abort() 阻止调用下一个中间件
    c.JSON() 返回JSON格式的响应
    c.String() 返回字符串格式的响应
    c.XML() 返回XML格式的响应
    c.YAML() 返回YAML格式的响应
    c.HTML() 返回HTML格式的响应
    c.File() 返回文件格式的响应
    c.Redirect() 返回重定向格式的响应
    
    c.Writer.Header().Set("Access-Control-Allow-Origin", "*")//跨域设置
    rid := c.GetHeader(HeaderXRequestID)//获取请求头Header中的rid
    c.Writer.Header().Set(HeaderXRequestID, rid)//设置响应头Header中的rid
    c.Set(HeaderXRequestID, rid)//设置上下文Context中的rid
    c.Next()

web.OK(c, gin.H{"data": h.service.Ping()}) //gin.H 其实是 map[string]interface{} 的类型别名，用来构造 JSON

r := gin.New()//构造一个gin引擎, 它实现了http.Handler接口,并构建了路由, 可以直接用在http.ListenAndServe中使用
r.Use(middleware.Recover())//中间件,相当于拦截器链
r.GET("healthz", func(c *gin.Context) {//构造一个get方法
    c.String(http.StatusOK, "ok")
})
v1 := r.Group("/api/v1")//构造一个组, 所有的路由都需要添加这个组的前缀

//处理路由和处理类的绑定
func (h *Handler) RegisterRoutes(routerGroup *gin.RouterGroup) {
	g := routerGroup.Group("/hello")
	g.GET("/ping", h.ping)
	g.GET("", h.greet) //api/v1/hello?name=Steve
}

//等待信号量, ctrl+c / kill
quit := make(chan os.Signal, 1)                      // 创建一个容量为1的信号通道，用来接收系统发来的中断信号
signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 注册监听两种信号：SIGINT（Ctrl+C）和 SIGTERM（kill命令默认信号）
<-quit                                               // 阻塞等待信号到来（程序会在这里卡住，直到收到信号）
_ = app.Stop(srv, 10*time.Second)                    // 收到信号后开始执行优雅关闭流程，最多等待10秒
fmt.Println("app stop")
```