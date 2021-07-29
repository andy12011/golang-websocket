package game

import (
	"fmt"
	"websocket/services/responseservice"
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

func AddToAllPlayer(player *Player) {
	allConnectRecords.Mutex.Lock()
	defer allConnectRecords.Mutex.Unlock()

	if prePlayer, ok := allConnectRecords.Players[player.Token]; ok { //reconnection
		prePlayer.Mutex.Lock()
		defer prePlayer.Mutex.Unlock()

		prePlayer.PushJson(responseservice.GetResponse(responseservice.PLAYER_IS_CONNETION_ELSEWHERE, nil))

		newConn := player.Conn
		player.CopyPlayer(prePlayer)
		player.Conn = newConn
		player.IsReconnect = true
		prePlayer.Conn.Close()
	}

	allConnectRecords.Players[player.Token] = player
}

func RemoveFromAllPlayer(token string) {
	allConnectRecords.Mutex.Lock()
	defer allConnectRecords.Mutex.Unlock()

	player, ok := allConnectRecords.Players[token]
	if player.IsReconnect {
		player.IsReconnect = false
		return
	}

	if ok && !player.IsPlaying {
		player.LeaveRoom()
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
	utils.PrintWithTimeStamp(fmt.Sprintf("目前開啟房間數量: %d", len(AllRoom.Rooms)))

	for _, room := range AllRoom.Rooms {
		msg := fmt.Sprintf("房間編號: %s 玩家人數: %d", room.GetRoomUUID(), len(room.ConnectRecord.Players))
		utils.PrintWithTimeStamp(msg)

		for _, player := range room.ConnectRecord.Players {
			var msg string
			if player.Token == room.Owner.Token {
				msg = fmt.Sprintf("房間編號: %s 房長玩家:%s", room.GetRoomUUID(), player.Token)
			} else {
				msg = fmt.Sprintf("房間編號: %s 玩家:%s", room.GetRoomUUID(), player.Token)
			}
			utils.PrintWithTimeStamp(msg)
		}
	}
}
