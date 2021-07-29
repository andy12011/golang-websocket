package controllers

import (
	"fmt"
	"log"
	"net/http"
	"websocket/services/httpservice"
	"websocket/services/websocketservice"

	"github.com/google/uuid"
)

func WebsocketConn(writer http.ResponseWriter, request *http.Request) {
	token, _ := httpservice.GetUrlQuery("token", request)

	validate := websocketservice.ValidateToken(token)

	if !validate {
		log.Println("validateToken fail")
		return
	}

	websocketservice.WsUpgrade(writer, request, token)
}

func newToken() string {
	return fmt.Sprintf("%s-%s","player-", uuid.New().String())
}
