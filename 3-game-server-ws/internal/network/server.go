package network

import (
	"fmt"
	"game-server-ws/internal/config"
	"game-server-ws/internal/game"
	"game-server-ws/internal/protocol"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true //允许跨域
	},
}

type GameServer struct {
	room *game.Room
}

func NewGameServer() *GameServer {
	return &GameServer{
		room: game.NewRoom(),
	}
}

// handlerWebsocket 方法用于处理客户端的 WebSocket 连接请求
// w 是 http.ResponseWriter 类型，用于向客户端发送 HTTP 响应
// r 是 *http.Request 类型，包含了客户端的 HTTP 请求信息
func (s *GameServer) handlerWebsocket(w http.ResponseWriter, r *http.Request) {
	// 使用 upgrade.Upgrade 函数将 HTTP 连接升级为 WebSocket 连接
	// w 是响应写入器，r 是请求对象，nil 表示使用默认的响应头
	conn, err := upgrade.Upgrade(w, r, nil)
	// 检查升级过程中是否出错
	if err != nil {
		// 若出错，打印错误信息
		log.Println("failed to upgrade:", err)
		// 终止当前函数的执行
		return
	}
	// 调用 lit-game.NewPlayer 函数创建一个新的玩家对象
	// 将升级后的 WebSocket 连接 conn 赋值给玩家对象
	player := game.NewPlayer(conn) //把连接赋值给player

	//new出来就开始心跳 服务器每隔多少秒向客户端推送一个心跳包
	player.StartHeartbeat()

	// 调用 s.room.Join 方法将新创建的玩家加入游戏房间
	s.room.Join(player) //把玩家加入房间

	// 启动一个新的 goroutine 来执行 s.readLoop 方法
	// 该方法用于持续读取玩家发送的消息
	go s.readLoop(player)
}

// readLoop 方法用于持续从玩家的 WebSocket 连接中读取消息
// player 是 *game.Player 类型，代表要读取消息的玩家对象
func (s *GameServer) readLoop(player *game.Player) {
	// 使用 defer 关键字定义一个匿名函数，该函数会在 readLoop 方法返回时执行
	// 作用是将玩家从房间中移除，并关闭玩家的 WebSocket 连接
	defer func() {
		s.room.Leave(player)
		player.Conn.Close()
	}()

	// 进入无限循环，持续读取消息
	for {
		// 声明一个 protocol.Message 类型的变量 msgs，用于存储读取到的消息
		var msg protocol.Message
		// 从玩家的 WebSocket 连接中读取 JSON 格式的消息，并将其解析到 msgs 变量中
		err := player.Conn.ReadJSON(&msg)
		// 检查读取过程中是否出错
		if err != nil {
			// 若出错，打印错误信息
			log.Println("read error:", err)
			// 退出 readLoop 方法，触发 defer 函数
			return
		}
		// 调用 s.room.HandlerMessage 方法，将玩家和读取到的消息传递给游戏房间进行处理
		s.room.HandlerMessage(player, msg)
	}
}

/*
*
  - 1.启动一个httpServer ListenAndServer
  - 2. 当/ws 客户端建立的时候， 路由到handlerWebsocket方法里
    2-1： 创建一个Player
    2-2: Player开启心跳
    2-3： Player加入Room
    2-4： 开启一个goroutine， 持续读取客户端发送的消息
*/
func (s *GameServer) Start() error {
	http.HandleFunc("/ws", s.handlerWebsocket) //这个方法只有在客户端连接的时候才会调用  比如：ws://localhost:8888/ws  /ws就会路由到这个方法里
	addr := fmt.Sprintf("%s:%d", config.GlobalConfig.Server.Host, config.GlobalConfig.Server.Port)
	log.Println("start server at", addr)
	server := &http.Server{
		Addr:         addr,
		ReadTimeout:  config.GlobalConfig.Server.ReadTimeout,
		WriteTimeout: config.GlobalConfig.Server.WriteTimeout,
		IdleTimeout:  config.GlobalConfig.Server.IdeleTimeout,
	}
	return server.ListenAndServe()
}
