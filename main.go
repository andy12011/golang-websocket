package main

import (
	"fmt"
	"net/http"
	"time"
	"websocket/config"
	"websocket/database/redis"
	"websocket/game"
	"websocket/oauth2/google"
	"websocket/routes"
	"websocket/utils"
)

func main() {

	config.Init(".env")
	google.Init()
	routes.Init()
	redis.Connection()

	fmt.Println("server start at :8099")

	go showCurrentPlayerCount()

	http.ListenAndServe(":8099", nil)
}

func showCurrentPlayerCount() {
	for {
		msg := fmt.Sprintf("線上玩家數量: %d", game.GetOnlinePlayerCount())
		utils.PrintWithTimeStamp(msg)
		game.ShowAllRoomInfo()
		time.Sleep(4 * time.Second)
	}
}
