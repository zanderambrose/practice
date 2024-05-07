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

func PostVenueData(url string, postable interface{}) *http.Response {
	venue := strings.ToLower(url)
	payload, err := json.Marshal(postable)
	if err != nil {
		// TODO - Log handling
		fmt.Println("error on that marshal mathers", err)
	}
	// TODO - This needs env variable
	resp, err := http.Post(fmt.Sprintf("http://server:8080/api/v1/%s", venue), "application/json", bytes.NewBuffer(payload))

	if err != nil {
		// TODO - Log handling
		fmt.Println("error on that http req", err)
	}

	return resp
}

func GetCurrentTime() string {
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		// TODO - Log handling
		fmt.Println("Error loading location:", err)
	}
	currentTime := time.Now().In(loc)
	return currentTime.Format("2006-01-02 15:04:05")
}
