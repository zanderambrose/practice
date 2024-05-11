package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
	"whoshittin/scraper/venueNames"
)

type NormalizedEventTime struct {
	Start string
	End   string
}

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

const STANDARD_DATE_LAYOUT = "2006-01-02"
const STANDARD_TIME_LAYOUT = "3:04 PM"

func formatTimeString(time time.Time) string {
	return time.Format(STANDARD_TIME_LAYOUT)
}

func parseTimeString(timeStr string) (time.Time, error) {
	return time.Parse(STANDARD_TIME_LAYOUT, timeStr)
}

func NormalizeTime(timeString string) (string, string, error) {
	times := strings.Split(timeString, "-")
	startTimeStr := strings.Trim(times[0], " ")
	endTimeStr := strings.Trim(times[1], " ")
	startTime, err := parseTimeString(startTimeStr)
	if err != nil {
		fmt.Println("Error parsing start time:", err)
		return "", "", err
	}
	endTime, err := parseTimeString(endTimeStr)
	if err != nil {
		fmt.Println("Error parsing end time:", err)
		return "", "", err
	}
	return formatTimeString(startTime), formatTimeString(endTime), nil
}

func NormalizeTimes(timeString string) ([]NormalizedEventTime, error) {
	var eventTimes []NormalizedEventTime
	times := strings.Split(timeString, "&")
	for i := 0; i < len(times); i++ {
		var eventTime NormalizedEventTime
		startTimeStr := strings.Trim(times[i], " ")
		parsedStartTime, err := parseTimeString(startTimeStr)
		if err != nil {
			fmt.Println("Error parsing start time:", err)
			return eventTimes, errors.New(err.Error())
		}
		parsedEndTime := parsedStartTime.Add(time.Hour + 15*time.Minute)
		eventTime.Start = formatTimeString(parsedStartTime)
		eventTime.End = formatTimeString(parsedEndTime)
		eventTimes = append(eventTimes, eventTime)
	}
	return eventTimes, nil
}

func NormalizeDate(dateString string, venue string) (time.Time, error) {
	currentYear := strconv.Itoa(time.Now().Year())
	venueDateFormats := map[string]string{
		"ornithology":  "Monday, January 2, 2006",
		"ornithology2": "Mon, January 2, 2006",
		"ornithology3": "Mon, Apr 2, 2006",
		"smalls":       "Mon Jan 02 2006",
		"mezzrow":      "Mon Jan 02 2006",
	}

	var parsedDate time.Time
	var err error

	if venue == venueNames.Smalls || venue == venueNames.Mezzrow {
		parsedDate, err = time.Parse(venueDateFormats[venue], dateString+" "+currentYear)

	}
	if venue == venueNames.OrnithologyVenueName {
		parsedDate, err = time.Parse(venueDateFormats[venue], dateString)
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
	}

	normalizedDateStr := parsedDate.Format(STANDARD_DATE_LAYOUT)
	normalizedDate, err := time.Parse(STANDARD_DATE_LAYOUT, normalizedDateStr)
	if err != nil {
		return time.Time{}, err
	}

	return normalizedDate, nil
}
