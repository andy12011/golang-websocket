package responseservice

type response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

const (
	SUCCESS             = 20001
	CREATE_ROOM_SUCCESS = 20002

	PLAYER_IS_NOT_IN_ROOM = 41001
	PLAYER_IS_IN_ROOM     = 41002

	ROOM_UUID_ERROR     = 41201
	ROOM_UUID_KEY_ERROR = 41202

	PARAMS_ERROR = 42201
	EVENT_ERROR  = 42202
)

var messages = map[int]string{
	SUCCESS:             "Success.",
	CREATE_ROOM_SUCCESS: "Create room success.",

	PARAMS_ERROR: "Params Error.",
	EVENT_ERROR:  "This event is wrong.",

	ROOM_UUID_ERROR:     "Not fount room num.",
	ROOM_UUID_KEY_ERROR: "Not fount correct key.",

	PLAYER_IS_NOT_IN_ROOM: "Player is not in any room.",
	PLAYER_IS_IN_ROOM:     "Player is in a room.",
}
