package redis

import (
	"testing"
	"websocket/config"
)

func TestConnection(t *testing.T) {
	config.Init("../../.env")
	Connection()
}
