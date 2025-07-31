package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Request-ID")
		/**
		浏览器在发送真正的跨域请求之前，会先自动发一个 OPTIONS 请求，询问服务器“我能不能用这些方法和头？”
		c.AbortWithStatus(204)
		服务器收到后，直接返回 204 No Content（空响应体），告诉浏览器“可以，继续发正式请求吧”。
		*/
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
