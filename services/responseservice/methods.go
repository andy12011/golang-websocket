package responseservice

func GetResponse(code int, data interface{}) *response {
	return &response{
		Code: code,
		Msg: getMsg(code),
		Data: data,
	}
}

func GetSuccessResponse() *response {
	return &response{
		Code: SUCCESS,
		Msg: getMsg(SUCCESS),
	}
}

func GetEventErrorResponse() *response {
	return &response{
		Code: EVENT_ERROR,
		Msg: getMsg(EVENT_ERROR),
	}
}

func getMsg(code int) string {
	if msg, ok := messages[code]; ok {
		return msg
	}

	return ""
}