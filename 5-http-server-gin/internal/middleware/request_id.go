package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const HeaderXRequestID = "X-Request-ID"

// RequestID 从请求头中获取 X-Request-ID，若不存在则生成一个新的 UUID。
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		rid := c.GetHeader(HeaderXRequestID)
		if rid == "" {
			rid = uuid.New().String()
		}
		c.Writer.Header().Set(HeaderXRequestID, rid)
		c.Set(HeaderXRequestID, rid)
		c.Next()
	}
}
