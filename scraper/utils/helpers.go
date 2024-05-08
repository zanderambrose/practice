package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
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

const standardLayout = "2006-01-02"

func NormalizeDate(dateString string, venue string) (time.Time, error) {
	venueDateFormats := map[string]string{
		"ornithology":  "Monday, January 2, 2006",
		"ornithology2": "Mon, January 2, 2006",
		"ornithology3": "Mon, Apr 2, 2006",
	}

	parsedDate, err := time.Parse(venueDateFormats[venue], dateString)
	if err != nil {
		parsedDate, err = time.Parse(venueDateFormats[venue+"2"], dateString)
		if err != nil {
			parsedDate, err = time.Parse(venueDateFormats[venue+"3"], dateString)
			if err != nil {
				fmt.Println("Error formatting ", venue, " data.")
				return time.Time{}, errors.New("Error formatting venue data")
			}
		}
	}

	normalizedDateStr := parsedDate.Format(standardLayout)
	fmt.Println("normalizedDateStr: ", normalizedDateStr)
	normalizedDate, err := time.Parse(standardLayout, normalizedDateStr)
	fmt.Println("normalizedDate: ", normalizedDate)
	if err != nil {
		return time.Time{}, err
	}

	return normalizedDate, nil
}
