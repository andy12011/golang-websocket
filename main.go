package main

import (
	"fmt"
	"net/http"
	"time"
	"websocket/game"
	"websocket/routes"
	"websocket/utils"
)

func main() {
	routes.Init()
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
