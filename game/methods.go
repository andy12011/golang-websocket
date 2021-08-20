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
	return &ConnectRecord{}
}

func AddToAllPlayer(player *Player) {

	if val, ok := allConnectRecords.Players.Load(player.Token); ok { //reconnection
		prePlayer, _ := val.(*Player)

		prePlayer.PushJson(responseservice.GetResponse(responseservice.PLAYER_IS_CONNETION_ELSEWHERE, nil))

		newConn := player.Conn
		player.CopyPlayer(prePlayer)
		player.Conn = newConn
		player.IsReconnect = true
		prePlayer.Conn.Close()
	}

	allConnectRecords.Players.Store(player.Token, player)
}

func RemoveFromAllPlayer(token string) {

	val, ok := allConnectRecords.Players.Load(token)

	if !ok {
		return
	}

	player, _ := val.(*Player)

	if player.IsReconnect {
		player.IsReconnect = false
		return
	}

	player.LeaveRoom()
	allConnectRecords.Players.Delete(player.Token)
}

func GetOnlinePlayerCount() int {
	return utils.GetSyncMapLen(allConnectRecords.Players)
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
		msg := fmt.Sprintf("房間編號: %s 玩家人數: %d", room.GetRoomUUID(), utils.GetSyncMapLen(room.ConnectRecord.Players))
		utils.PrintWithTimeStamp(msg)

		allConnectRecords.Players.Range(func(key, value interface{}) bool {
			player, _ := value.(*Player)
			var msg string

			if player.Token == room.Owner.Token {
				msg = fmt.Sprintf("房間編號: %s 房長玩家: %s Token: %s", room.GetRoomUUID(), player.Nickname, player.Token)
			} else {
				msg = fmt.Sprintf("房間編號: %s 玩家: %s Token: %s", room.GetRoomUUID(), player.Nickname, player.Token)
			}
			utils.PrintWithTimeStamp(msg)
			return true
		})
	}
}
