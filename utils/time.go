package utils

import "time"

// yyyy-mm-dd hh:mm:ss
func ParseTimeTimeToString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}