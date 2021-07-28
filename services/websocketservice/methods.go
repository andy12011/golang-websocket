package websocketservice

import (
	"log"
	"net/http"
	"websocket/definitions"
	"websocket/services/responseservice"

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
		log.Println("upgrade:", err)
		return
	}

	player := definitions.NewPlayer()
	player.Conn = websocketConn
	player.Token = token

	player.AddToAllPlayer()

	defer func() {
		log.Println("disconnect !!")
		definitions.RemoveFromAllPlayer(player.Token)
		player.Conn.Close()
	}()

	player.PushJson(responseservice.GetSuccessResponse())

	for {
		_, msg, err := player.Conn.ReadMessage()

		if err != nil {
			log.Println("read:", err)
			break
		}

		player.ReceiveMsg(msg)
	}
}