package hello

import (
	"github.com/gin-gonic/gin"
	"http-server/internal/web"
)

type Handler struct {
	service Service //内嵌了一个接口
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) ping(c *gin.Context) {
	web.OK(c, gin.H{"data": h.service.Ping()}) //gin.H 其实是 map[string]interface{} 的类型别名，用来构造 JSON
}

func (h *Handler) greet(c *gin.Context) {
	name := c.Query("name")
	ms := h.service.Greet(name)
	web.OK(c, gin.H{"data": ms})
}

func (h *Handler) RegisterRoutes(routerGroup *gin.RouterGroup) {
	g := routerGroup.Group("/hello")
	g.GET("/ping", h.ping)
	g.GET("", h.greet) //api/v1/hello?name=Steve
}
