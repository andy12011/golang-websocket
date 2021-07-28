package definitions

import (
	"sync"
)

const (
	WS_CONN_TOKEN_PARAM = "token"
	TEXT_TYPE           = 1
)

var allConnectRecords = &ConnectRecord{
	ID:      "All",
	Players: make(map[string]*Player),
}

type ConnectRecord struct {
	ID      string
	Players map[string]*Player
	Mutex   sync.Mutex
}

type WebsocketEvent struct {
	Event string `json:"event"`
	Msg   string `json:"msg"`
}
