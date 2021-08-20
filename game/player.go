package game

import (
	"encoding/json"
	"fmt"
	"sync"
	"websocket/services/responseservice"

	"github.com/gorilla/websocket"
)

type Player struct {
	Nickname    string
	Token       string
	Conn        *websocket.Conn
	RoomUUID    string
	Position    int
	Character   string
	IsPlaying   bool
	IsReconnect bool
	Mutex       sync.Mutex
}

func (player *Player) AddToAllPlayer() {
	AddToAllPlayer(player)
}

func (player *Player) RemoveFromAllPlayer() {
	RemoveFromAllPlayer(player.Token)
}

func (player *Player) PushJson(v interface{}) {
	player.Conn.WriteJSON(v)
}

func (player *Player) Push(v interface{}) {
	player.Conn.WriteMessage(TEXT_TYPE, []byte(fmt.Sprintf("%s", v)))
}

func (player *Player) ReceiveMsg(msg []byte) {
	var event Event
	err := json.Unmarshal(msg, &event)

	if err != nil {
		player.PushJson(responseservice.GetEventErrorResponse())
		return
	}

	player.DoEvent(&event)

}

func (player *Player) DoEvent(event *Event) {
	player.Mutex.Lock()
	defer player.Mutex.Unlock()

	handler, ok := getEventHandler(event.Event)

	if !ok {
		player.PushJson(responseservice.GetEventErrorResponse())
		return
	}

	handler(player, event)
}

func (player *Player) EnterRoom(uuid string) {
	room, ok := GetRoom(player.RoomUUID)

	if ok {
		room.PlayerEnter(player)
	} else {
		player.PushJson(responseservice.GetResponse(responseservice.ROOM_UUID_ERROR, nil))
	}
}

func (player *Player) LeaveRoom() int {
	beforeRoomUUID := player.RoomUUID

	room, ok := GetRoom(beforeRoomUUID)

	if ok {
		room.PlayerLeave(player)
		return responseservice.SUCCESS
	} else {
		return responseservice.PLAYER_IS_NOT_IN_ROOM
	}
}

func (player *Player) SetPlayerRoomUUID(uuid string) {
	player.RoomUUID = uuid
}

func (player *Player) IsInRoom() bool {
	return player.RoomUUID != ""
}

func (player *Player) CopyPlayer(target *Player) {
	player.Nickname = target.Nickname
	player.Token = target.Token
	player.Conn = target.Conn
	player.RoomUUID = target.RoomUUID
	player.Position = target.Position
	player.Character = target.Character
	player.IsPlaying = target.IsPlaying
	player.IsReconnect = target.IsReconnect
}
