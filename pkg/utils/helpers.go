package utils

import (
	"context"
)

var CTX = context.Background()

func IsAdult(age int) bool {
	if age < 18 {
		return false
	}
	return true
}
