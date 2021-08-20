package utils

import (
	"fmt"
	"sync"
	"time"
)


func GetDateTimeString() string {
	return time.Now().Format("2006.01.02 15:04:05")
}

func PrintWithTimeStamp(msg string) {
	fmt.Printf("[%s] %s\n", GetDateTimeString(), msg)
}

func GetSyncMapLen(sMap sync.Map) (len int) {
	sMap.Range(func(key, value interface{}) bool {
		len++
		return true
	})

	return len
}