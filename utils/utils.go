package utils

import (
	"fmt"
	"time"
)


func GetDateTimeString() string {
	return time.Now().Format("2006.01.02 15:04:05")
}

func PrintWithTimeStamp(msg string) {
	fmt.Printf("[%s] %s", GetDateTimeString(), msg)
}