package game

import (
	"sync"
)

const (
	WS_CONN_TOKEN_PARAM = "token"
	TEXT_TYPE           = 1
)

var allConnectRecords = &ConnectRecord{
	ID:      "All",
}

type ConnectRecord struct {
	ID      string
	Players sync.Map
	Mutex   sync.Mutex
}

type WebsocketEvent struct {
	Event string `json:"event"`
	Msg   string `json:"msg"`
}
