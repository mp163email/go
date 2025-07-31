package app

import (
	"context"
	"net/http"
	"time"
)

/*
new一个http服务器

客户端请求流程：
建立连接 → 发送请求（受 ReadTimeout 限制）→
服务器处理 → 返回响应（受 WriteTimeout 限制）→
保持连接空闲（受 IdleTimeout 限制）→ 断开

防止慢速客户端攻击（Read）
避免服务端响应卡死（Write）
合理释放闲置资源（Idle）
*/
func NewHttpServer(h http.Handler, addr string) *http.Server {
	return &http.Server{
		Addr:         addr,
		Handler:      h,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}

func Start(s *http.Server) error {
	return s.ListenAndServe()
}

// 多久后关闭服务器 优雅关闭
func Stop(s *http.Server, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return s.Shutdown(ctx)
}
