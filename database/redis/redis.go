package redis

import (
	"fmt"
	"os"
	"time"
	"websocket/config"

	"github.com/go-redis/redis"
)

var redisClient *redis.Client

func Connection() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv(config.REDIS_HOST), os.Getenv(config.REDIS_PORT)),
		Password: os.Getenv(config.REDIS_PASSWORD),
		DB:       0,
	})
	
	SetDB(rdb)

	go listenRedisConnection()
}

func listenRedisConnection() {
	for {
		_, err := GetDB().Ping().Result()
		if nil != err {
			Connection()
			break
		}
		time.Sleep(1 * time.Minute)
	}
}

func SetDB(rdb *redis.Client) {
	redisClient = rdb
}

func GetDB() *redis.Client {
	return redisClient
}
