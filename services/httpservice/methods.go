package httpservice

import "net/http"

func GetUrlQuery(key string, request *http.Request) (string, bool) {
	keys, ok := request.URL.Query()[key]
	if !ok {
		return "", ok
	}
	return keys[0], ok
}