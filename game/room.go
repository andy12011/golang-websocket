package game

import (
	"sync"
	"websocket/services/responseservice"
	"websocket/utils"

	"github.com/google/uuid"
)

type allRoom struct {
	Rooms map[string]*Room
	Mutex sync.Mutex
}

var AllRoom = &allRoom{
	Rooms: make(map[string]*Room),
}

type Room struct {
	UUID          string
	Owner         *Player
	ConnectRecord *ConnectRecord
	IsStarting    bool
}

func (room *Room) Init(owner *Player) {
	room.SetRoomUUID(uuid.New().String())
	SetRoom(room)
	room.SetOwner(owner)
}

func (room *Room) SetRoomUUID(uuid string) {
	room.UUID = uuid
}

func (room *Room) GetRoomUUID() string {
	return room.UUID
}

func (room *Room) SetOwner(owner *Player) {
	room.Owner = owner
	room.PlayerEnter(owner)
}

func (room *Room) GetOwner() *Player {
	return room.Owner
}

func (room *Room) PlayerEnter(player *Player) {
	room.ConnectRecord.Mutex.Lock()
	defer room.ConnectRecord.Mutex.Unlock()

	player.SetPlayerRoomUUID(room.GetRoomUUID())
	room.ConnectRecord.Players[player.Token] = player
}

func (room *Room) PlayerLeave(player *Player) {
	room.ConnectRecord.Mutex.Lock()
	defer room.ConnectRecord.Mutex.Unlock()

	_, ok := room.ConnectRecord.Players[player.Token]

	if ok {
		delete(room.ConnectRecord.Players, player.Token)

		if len(room.ConnectRecord.Players) == 0 {
			CloseRoom(room)
		} else if player.Token == room.Owner.Token {
			for _, newOwner := range room.ConnectRecord.Players {
				room.Owner = newOwner
				break
			}
		}
	}

	player.SetPlayerRoomUUID("")
}

func (room *Room) Echo(sender *Player, msg string) {
	sendMsg := &ResponseMessage{
		CreatedAt: utils.GetDateTimeString(),
		Msg:       msg,
		Sender:    sender.Name,
	}
	for _, player := range room.ConnectRecord.Players {
		player.PushJson(responseservice.GetResponse(responseservice.SUCCESS, sendMsg))
	}
}
