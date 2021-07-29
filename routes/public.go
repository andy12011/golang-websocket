package routes

import (
	"fmt"
	"net/http"
	"websocket/controllers"
)


func Init() {
	http.HandleFunc("/ws/connection", controllers.WebsocketConn)

	http.HandleFunc("/google/oauth2", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello")
	})
}