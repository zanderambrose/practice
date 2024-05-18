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

func PostVenueData(venueName string, postable interface{}) *http.Response {
	venue := strings.ToLower(venueName)
	// TODO - This needs env variable
	url := fmt.Sprintf("http://server:8080/api/v1/%s", venue)
	payload, err := json.Marshal(postable)
	if err != nil {
		// TODO - Log handling
		fmt.Println("error on that marshal mathers", err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
	}
	req.Header.Set("Content-Type", "application/json")
	// TODO - This needs env variable
	req.Header.Set("Authorization", "Bearer admin")
	// TODO - This needs env variable
	req.Header.Set("X-Client-ID", "admin")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
	}
	defer resp.Body.Close()

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
const STANDARD_DATE_REPRESENTATION_LAYOUT = "Mon Jan 02 2006"
const STANDARD_TIME_LAYOUT = "3:04 PM"
const STANDARD_TIME_LAYOUT2 = "3:04PM"

func FormatTimeString(time time.Time) string {
	return time.Format(STANDARD_TIME_LAYOUT)
}

func ParseTimeString(timeStr string) (time.Time, error) {
	if strings.Contains(timeStr, " ") {
		return time.Parse(STANDARD_TIME_LAYOUT, timeStr)
	}
	return time.Parse(STANDARD_TIME_LAYOUT2, timeStr)
}

func NormalizeTime(timeString string) (string, string, error) {
	times := strings.Split(timeString, "-")
	startTimeStr := strings.Trim(times[0], " ")
	endTimeStr := strings.Trim(times[1], " ")
	startTime, err := ParseTimeString(startTimeStr)
	if err != nil {
		fmt.Println("Error parsing start time:", err)
		return "", "", err
	}
	endTime, err := ParseTimeString(endTimeStr)
	if err != nil {
		fmt.Println("Error parsing end time:", err)
		return "", "", err
	}
	return FormatTimeString(startTime), FormatTimeString(endTime), nil
}

func NormalizeTimes(timeString string) ([]NormalizedEventTime, error) {
	var eventTimes []NormalizedEventTime
	times := strings.Split(timeString, "&")
	for i := 0; i < len(times); i++ {
		var eventTime NormalizedEventTime
		startTimeStr := strings.Trim(times[i], " ")
		parsedStartTime, err := ParseTimeString(startTimeStr)
		if err != nil {
			fmt.Println("Error parsing start time:", err)
			return eventTimes, errors.New(err.Error())
		}
		parsedEndTime := parsedStartTime.Add(time.Hour + 15*time.Minute)
		eventTime.Start = FormatTimeString(parsedStartTime)
		eventTime.End = FormatTimeString(parsedEndTime)
		eventTimes = append(eventTimes, eventTime)
	}
	return eventTimes, nil
}

func NormalizeDate(dateString string, venue string) (time.Time, error) {
	currentYear := strconv.Itoa(time.Now().Year())
	venueDateFormats := map[string]string{
		venueNames.Ornithology:       "Monday, January 2, 2006",
		venueNames.Ornithology + "2": "Mon, January 2, 2006",
		venueNames.Ornithology + "3": "Mon, Apr 2, 2006",
		venueNames.Smalls:            "Mon Jan 02 2006",
		venueNames.Mezzrow:           "Mon Jan 02 2006",
		venueNames.Django:            "Monday, January 2 2006",
		venueNames.Smoke:             "Mon, Jan 2 2006",
		venueNames.BlueNote:          STANDARD_DATE_LAYOUT,
		venueNames.Zinc:              "Mon Jan 02 2006",
	}

	var parsedDate time.Time
	var err error

	if venue == venueNames.BlueNote {
		var formattedDateString string
		if len(dateString) == 1 {
			formattedDateString = "0" + dateString
		} else {
			formattedDateString = dateString
		}
		today := time.Now()
		todayFormatted := today.Format(STANDARD_DATE_LAYOUT)
		tokens := strings.Split(todayFormatted, "-")
		tokens[2] = formattedDateString
		fullDate := strings.Join(tokens, "-")
		parsedDate, err = time.Parse(venueDateFormats[venue], fullDate)
	}

	if venue == venueNames.Django {
		parsedDate, err = time.Parse(venueDateFormats[venue], dateString+" "+currentYear)
	}

	if venue == venueNames.Smoke {
		parsedDate, err = time.Parse(venueDateFormats[venue], dateString+" "+currentYear)
	}

	if venue == venueNames.Smalls || venue == venueNames.Mezzrow || venue == venueNames.Zinc {
		parsedDate, err = time.Parse(venueDateFormats[venue], dateString+" "+currentYear)

	}
	if venue == venueNames.Ornithology {
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

	if venue == venueNames.Vanguard {
		return time.Time{}, errors.New("Vanguard dates are handled in own scope")
	}

	normalizedDateStr := parsedDate.Format(STANDARD_DATE_LAYOUT)
	normalizedDate, err := time.Parse(STANDARD_DATE_LAYOUT, normalizedDateStr)
	if err != nil {
		return time.Time{}, err
	}

	return normalizedDate, nil
}

func ParseDate(dateStr string) (time.Month, int) {
	parts := strings.Split(dateStr, " ")
	if len(parts) != 2 {
		return time.January, 1 // Default values
	}

	monthStr := strings.ToUpper(parts[0])
	month, _ := time.Parse("January", monthStr) // Parse the month
	day := 1                                    // Default day
	fmt.Sscanf(parts[1], "%d", &day)            // Parse the day
	return month.Month(), day
}
