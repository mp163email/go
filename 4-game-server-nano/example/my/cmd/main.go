package main

import (
	"game-server-nano/example/my/comps"
	"github.com/lonng/nano"
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/pipeline"
	"github.com/lonng/nano/serialize/json"
	"log"
	"net/http"
	"strings"
)

func main() {
	components := &component.Components{} //注意：这里的关键字是Components 末尾有一个小写的s
	components.Register(
		comps.NewRoomManager(),
		component.WithName("room"),
		component.WithNameFunc(strings.ToLower),
	)

	//处理pipeline,添加自定义处理器
	pip := pipeline.New()
	var stats = &comps.Status{}
	pip.Outbound().PushBack(stats.Outbound)
	pip.Inbound().PushFront(stats.Inbound)

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	
	nano.Listen(":3250",
		nano.WithPipeline(pip),
		nano.WithComponents(components),
		nano.WithDebugMode(),
		nano.WithIsWebsocket(true),
		nano.WithSerializer(json.NewSerializer()),
		nano.WithCheckOriginFunc(func(r *http.Request) bool {
			return true
		}),
		nano.WithWSPath("/nano"),
	)
}
