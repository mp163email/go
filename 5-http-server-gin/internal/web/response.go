package web

import "github.com/gin-gonic/gin"

type Resp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func OK(c *gin.Context, data interface{}) {
	c.JSON(200, Resp{Code: 0, Message: "ok", Data: data})
}

func Fail(c *gin.Context, httpCode, bizCode int, message string) {
	c.JSON(httpCode, Resp{Code: bizCode, Message: message})
}
