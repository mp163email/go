package middleware

import (
	"github.com/gin-gonic/gin"
	"http-server/internal/web"
	"net/http"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			//在发生 panic 时，如果某个 defer 里调用了 recover()，就能把 panic 的值拿到手，从而阻止整个程序崩溃，并允许你优雅地记录日志、返回错误响应等。
			if rec := recover(); rec != nil {
				web.Fail(c, http.StatusInternalServerError, web.CodeServerError, "internal server error")
				c.Abort()
			}
		}()
	}
}
