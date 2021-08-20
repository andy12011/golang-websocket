package httpservice

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"time"
	"websocket/config"
	"websocket/utils"

	"github.com/google/uuid"
)

func GetUrlQuery(key string, request *http.Request) (string, bool) {
	keys, ok := request.URL.Query()[key]
	if !ok {
		return "", ok
	}
	return keys[0], ok
}

func NewAuthCookie() *http.Cookie {
	cookie := base64.URLEncoding.EncodeToString([]byte(uuid.New().String()))
	expires := time.Now().Add(10 * time.Minute)

	msg := fmt.Sprintf("建立 Cookie 成功: %s 到期時間: %s", cookie, utils.ParseTimeTimeToString(expires))
	utils.PrintWithTimeStamp(msg)

	return &http.Cookie{
		Name:    os.Getenv(config.APP_NAME),
		Value:   cookie,
		Expires: expires,
	}
}

func GetAuthCookie(r *http.Request) (c string, hasAuthCookie bool) {
	cookie, err := r.Cookie(os.Getenv(config.APP_NAME))

	if err != nil {
		return c, hasAuthCookie
	}
	c = cookie.Value
	hasAuthCookie = true

	msg := fmt.Sprintf("取得 Cookie 成功: %s 到期時間: %s", c, utils.ParseTimeTimeToString(cookie.Expires))
	utils.PrintWithTimeStamp(msg)

	return c, hasAuthCookie
}
