package main

import (
	localcomp "game-server-nano/internal/component"
	"game-server-nano/internal/service"
	"github.com/lonng/nano"
	nanocomp "github.com/lonng/nano/component"
	"github.com/lonng/nano/pipeline"
	"github.com/lonng/nano/serialize/json"
	"github.com/lonng/nano/session"
	"log"
	"net/http"
	"strings"
)

func messageHandler(s *session.Session, msg *pipeline.Message) error {
	log.Printf("Received message: %+v", msg)
	return nil
}

func main() {
	//初始化核心组件
	components := &nanocomp.Components{}
	components.Register(
		service.NewRoomService(localcomp.NewRoommanager()), //注册服务组件
		nanocomp.WithName("room"),
		nanocomp.WithNameFunc(strings.ToLower),
	)
	pip := pipeline.New()
	pip.Inbound().PushBack(messageHandler)

	nano.Listen(":3250",
		nano.WithPipeline(pip),
		nano.WithComponents(components),
		nano.WithIsWebsocket(true),
		nano.WithSerializer(json.NewSerializer()),
		nano.WithCheckOriginFunc(func(r *http.Request) bool {
			return true
		}),
		nano.WithDebugMode(),
		nano.WithWSPath("/"),
		//nano.WithWSPath("/nano"),
	)

}
