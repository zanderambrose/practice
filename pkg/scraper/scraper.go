package scraper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Scraper() {
	c := colly.NewCollector()

	allPerformers := make([]map[string]string, 0)

	c.OnHTML("div.artists", func(e *colly.HTMLElement) {
		performers := make(map[string]string)

		e.ForEach("a.title5", func(_ int, elem *colly.HTMLElement) {
			text := elem.Text
			parts := strings.Split(text, " / ")
			if len(parts) == 2 {
				performer := strings.TrimSpace(parts[0])
				instrument := strings.TrimSpace(parts[1])
				performers[performer] = instrument
			}
		})

		allPerformers = append(allPerformers, performers)
	})

	c.Visit("https://www.smallslive.com/")

	for i, performers := range allPerformers {
		isEarlySet := isEarlySet(i)
		currentTime := time.Now()
		venue := determineVenue(i)
		performers["isEarlySet"] = strconv.FormatBool(isEarlySet)
		performers["currentTime"] = currentTime.Format("2006-01-02 15:04:05")
		performers["venue"] = venue
		payload, err := json.Marshal(performers)
		if err != nil {
			fmt.Println("error on that marshal mathers", err)
		}
		resp, err := http.Post(fmt.Sprintf("http://localhost:8080/api/v1/%s", venue), "application/json", bytes.NewBuffer(payload))

		if err != nil {
			fmt.Println("error on that http req", err)
		}

		fmt.Println("resp: ", resp)
	}

}

func isEarlySet(index int) bool {
	if index > 1 {
		return false
	}
	return true
}

func determineVenue(index int) string {
	if index%2 == 0 {
		return "smalls"
	}
	return "mezzrow"
}
