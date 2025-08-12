package main

import (
	"game-server-ws/internal/config"
	"game-server-ws/internal/network"
	"log"
)

func main() {
	_, err := config.Load("config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config:%v", err)
	}

	gameServer := network.NewGameServer()

	if err := gameServer.Start(); err != nil {
		log.Fatalf("Failed to start lit-game server:%v", err)
	}
}
