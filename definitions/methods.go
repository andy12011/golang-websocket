package definitions

import (
	"fmt"
	"websocket/utils"
)

func NewPlayer() *Player {
	return &Player{}
}

func NewWsEvent() *WebsocketEvent {
	return &WebsocketEvent{}
}

func NewRoom() *Room {
	return &Room{
		ConnectRecord: NewWsConnetions(),
	}
}

func NewWsConnetions() *ConnectRecord {
	return &ConnectRecord{
		Players: make(map[string]*Player),
	}
}

func AddToAllPlayer(ws *Player) {
	allConnectRecords.Mutex.Lock()
	defer allConnectRecords.Mutex.Unlock()
	allConnectRecords.Players[ws.Token] = ws
}

func RemoveFromAllPlayer(token string) {
	allConnectRecords.Mutex.Lock()
	defer allConnectRecords.Mutex.Unlock()

	_, ok := allConnectRecords.Players[token]
	if ok {
		delete(allConnectRecords.Players, token)
	}
}

func GetOnlinePlayerCount() int {
	allConnectRecords.Mutex.Lock()
	defer allConnectRecords.Mutex.Unlock()

	return len(allConnectRecords.Players)
}

func GetRoomCount() int {
	return len(AllRoom.Rooms)
}

func GetRoom(roomUUID string) (room *Room, ok bool) {
	AllRoom.Mutex.Lock()
	defer AllRoom.Mutex.Unlock()

	room, ok = AllRoom.Rooms[roomUUID]

	return room, ok
}

func SetRoom(room *Room) {
	AllRoom.Mutex.Lock()
	defer AllRoom.Mutex.Unlock()

	AllRoom.Rooms[room.GetRoomUUID()] = room
}

func CloseRoom(room *Room) {
	AllRoom.Mutex.Lock()
	defer AllRoom.Mutex.Unlock()

	delete(AllRoom.Rooms, room.UUID)
}

func ShowAllRoomInfo() {
	utils.PrintWithTimeStamp(fmt.Sprintf("目前開啟房間數量: %d\n", len(AllRoom.Rooms)))

	for _, room := range AllRoom.Rooms {
		msg := fmt.Sprintf("房間編號: %s 玩家人數: %d\n", room.GetRoomUUID(), len(room.ConnectRecord.Players))
		utils.PrintWithTimeStamp(msg)

		for _, player := range room.ConnectRecord.Players {
			var msg string
			if player.Token == room.Owner.Token {
				msg = fmt.Sprintf("房間編號: %s 房長玩家:%s\n", room.GetRoomUUID(), player.Token)
			} else {
				msg = fmt.Sprintf("房間編號: %s 玩家:%s\n", room.GetRoomUUID(), player.Token)
			}
			utils.PrintWithTimeStamp(msg)
		}
	}
}
