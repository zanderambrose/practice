package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
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

func PostVenueData(url string, postable interface{}) {
	venue := strings.ToLower(url)
	payload, err := json.Marshal(postable)
	if err != nil {
		fmt.Println("error on that marshal mathers", err)
	}
	resp, err := http.Post(fmt.Sprintf("http://server:8080/api/v1/%s", venue), "application/json", bytes.NewBuffer(payload))

	if err != nil {
		fmt.Println("error on that http req", err)
	}

	fmt.Println("resp: ", resp)
}
