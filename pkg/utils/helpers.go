package utils

import (
	"context"
	"time"
)

var CTX = context.Background()

func IsAdult(age int) bool {
	if age < 18 {
		return false
	}
	return true
}

func AppendCurrentTime(postable *map[string]string) {
	currentTime := time.Now()
	(*postable)["currentTime"] = currentTime.Format("2006-01-02 15:04:05")
}
