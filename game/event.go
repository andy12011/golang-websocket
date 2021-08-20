package game

import (
	"websocket/services/responseservice"
)

const (
	CREATE_ROOM_EVENT          = "createRoom"
	ENTER_ROOM_EVENT           = "enterRoom"
	LEAVE_ROOM_EVENT           = "leaveRoom"
	SEND_MESSAGE_TO_ROOM_EVENT = "sendRoomMessage"
)

type Event struct {
	Event string      `json:"event" binding:"required"`
	Data  interface{} `json:"data,omitempty"`
}

type SendMessage struct {
	Msg string `json:"msg" binding:"required"`
}

type ResponseMessage struct {
	CreatedAt string `json:"created_at" binding:"required"`
	Msg       string `json:"msg" binding:"required"`
	Sender    string `json:"sender" binding:"required"`
	Nickname  string `json:"nickname" binding:"required"`
}

var eventHandler = map[string]func(player *Player, event *Event){
	CREATE_ROOM_EVENT:          eventCreateRoomHandler,
	ENTER_ROOM_EVENT:           eventEnterRoomHandler,
	LEAVE_ROOM_EVENT:           eventLeaveRoomHandler,
	SEND_MESSAGE_TO_ROOM_EVENT: eventSendRoomMessageHandler,
}

func getEventHandler(eventType string) (func(player *Player, event *Event), bool) {
	handler, ok := eventHandler[eventType]

	return handler, ok
}

func eventCreateRoomHandler(player *Player, event *Event) {
	if player.IsInRoom() {
		player.PushJson(responseservice.GetResponse(responseservice.PLAYER_IS_IN_ROOM, nil))
		return
	}

	room := NewRoom()
	room.Init(player)

	player.PushJson(responseservice.GetResponse(responseservice.CREATE_ROOM_SUCCESS, nil))
}

func eventEnterRoomHandler(player *Player, event *Event) {
	if player.IsInRoom() {
		player.PushJson(responseservice.GetResponse(responseservice.PLAYER_IS_IN_ROOM, nil))
		return
	}

	if data, ok := event.Data.(map[string]interface{}); ok {
		if v, ok := data["room_uuid"]; !ok {
			player.PushJson(responseservice.GetResponse(responseservice.ROOM_UUID_KEY_ERROR, nil))
			return
		} else {
			room_uuid, ok := v.(string)

			if !ok {
				player.PushJson(responseservice.GetResponse(responseservice.ROOM_UUID_ERROR, nil))
				return
			}
			room, ok := GetRoom(room_uuid)

			if ok {
				room.PlayerEnter(player)
				player.PushJson(responseservice.GetSuccessResponse())
			} else {
				player.PushJson(responseservice.GetResponse(responseservice.ROOM_UUID_ERROR, nil))
			}
			return
		}
	}

	player.PushJson(responseservice.GetResponse(responseservice.EVENT_ERROR, nil))
}

func eventLeaveRoomHandler(player *Player, event *Event) {
	code := player.LeaveRoom()

	player.PushJson(responseservice.GetResponse(code, nil))
}

func eventSendRoomMessageHandler(player *Player, event *Event) {
	if !player.IsInRoom() {
		player.PushJson(responseservice.GetResponse(responseservice.PLAYER_IS_NOT_IN_ROOM, nil))
		return
	}

	if data, ok := event.Data.(map[string]interface{}); ok {
		if v, ok := data["msg"]; !ok {
			player.PushJson(responseservice.GetResponse(responseservice.EVENT_ERROR, nil))
			return
		} else {
			msg, ok := v.(string)

			if !ok {
				player.PushJson(responseservice.GetResponse(responseservice.EVENT_ERROR, nil))
				return
			}

			room, ok := GetRoom(player.RoomUUID)

			if ok {
				room.Echo(player, msg)
			}

			return
		}
	}
}
