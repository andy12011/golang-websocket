package controllers

import (
	"log"
	"net/http"
	"websocket/definitions"
	"websocket/services/httpservice"
	"websocket/services/responseservice"
	"websocket/services/websocketservice"
)

func WebsocketConn(writer http.ResponseWriter, request *http.Request) {
	token, ok := httpservice.GetUrlQuery(definitions.WS_CONN_TOKEN_PARAM, request)
	if !ok {
		log.Println("getQuery fail", responseservice.GetResponse(responseservice.PARAMS_ERROR, nil))
		return
	}

	validate := websocketservice.ValidateToken(token)

	if !validate {
		log.Println("validateToken fail")
		return
	}

	websocketservice.WsUpgrade(writer, request, token)
}
