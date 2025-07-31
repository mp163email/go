package router

import (
	"github.com/gin-gonic/gin"
	"http-server/internal/hello"
	"http-server/internal/middleware"
	"net/http"
)

type Options struct {
	HelloHandler *hello.Handler
}

func New(o Options) *gin.Engine {
	r := gin.New()

	//这四行代码依次给 Gin 引擎 挂载了 4 个中间件（middleware），它们会在 每个请求真正到达业务 handler 之前 被执行一次，用来完成统一的预处理或兜底保护
	//一句话：它们按顺序层层包裹，为所有接口提供“防崩溃、统一错误格式、请求追踪、跨域支持”的通用能力
	r.Use(gin.Recovery())
	r.Use(middleware.Recover())
	r.Use(middleware.RequestID())
	r.Use(middleware.CORS())

	//健康检查与元信息 解释说和k8s有关呢?
	r.GET("healthz", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	r.GET("readyz", func(c *gin.Context) {
		c.String(http.StatusOK, "ready")
	})
	r.GET("version", func(c *gin.Context) {
		c.JSON(200, gin.H{"version": "v1.0.0"})
	})

	//配置路由
	v1 := r.Group("/api/v1")
	o.HelloHandler.RegisterRoutes(v1)

	return r
}
