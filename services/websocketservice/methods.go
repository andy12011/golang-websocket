package websocketservice

import (
	"net/http"
	"websocket/game"
	"websocket/services/responseservice"
	"websocket/utils"

	"github.com/gorilla/websocket"
)

func ValidateToken(token string) bool {
	return true
}

func WsUpgrade(writer http.ResponseWriter, request *http.Request, token string) {
	upgrader := &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	websocketConn, err := upgrader.Upgrade(writer, request, nil)

	if err != nil {
		utils.PrintWithTimeStamp(err.Error())
		return
	}

	player := game.NewPlayer()
	player.Conn = websocketConn
	player.Token = token

	player.AddToAllPlayer()

	defer func() {
		utils.PrintWithTimeStamp("disconnect !!")
		game.RemoveFromAllPlayer(player.Token)
		player.Conn.Close()
	}()

	player.PushJson(responseservice.GetSuccessResponse())

	for {
		_, msg, err := player.Conn.ReadMessage()

		if err != nil {
			utils.PrintWithTimeStamp(err.Error())
			break
		}

		player.ReceiveMsg(msg)
	}
}
