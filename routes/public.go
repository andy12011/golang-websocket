package routes

import (
	"net/http"
	"websocket/controllers"
)


func Init() {
	http.HandleFunc("/ws/connection", controllers.WebsocketConn)
}