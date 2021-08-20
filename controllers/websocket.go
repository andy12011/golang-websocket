package controllers

import (
	"fmt"
	"net/http"
	"websocket/services/httpservice"
	"websocket/services/websocketservice"

	"github.com/google/uuid"
)

func WebsocketConn(writer http.ResponseWriter, request *http.Request) {
	nickname, _ := httpservice.GetUrlQuery("nickname", request)

	authCookie, hasAuthCookie := httpservice.GetAuthCookie(request)

	if !hasAuthCookie {
		cookie := httpservice.NewAuthCookie()
		authCookie = cookie.Value
		http.SetCookie(writer, cookie)
	}

	websocketservice.WsUpgrade(&writer, request, authCookie, nickname)
}

func newToken() string {
	return fmt.Sprintf("%s-%s", "player-", uuid.New().String())
}
