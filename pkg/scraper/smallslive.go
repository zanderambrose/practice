package scraper

import (
	"errors"
	"fmt"
	"github.com/gocolly/colly"
	"practice/pkg/utils"
	"strings"
)

func SmallsLiveScraper(c *colly.Collector) {
	c.OnHTML("article.event-display-today-and-tomorrow", func(e *colly.HTMLElement) {
		performers := make(map[string]string)

		ariaLabel := e.ChildAttr("a", "aria-label")
		eventDetails := strings.Split(ariaLabel, ", ")

		// Get event title
		eventTitle := e.ChildText("p.event-info-title")
		appendEventTitle(eventTitle, &performers)

		// Get venue info
		venueInfo, err := getEventDetails(eventDetails, -2)
		if err != nil {
			fmt.Println("Error getting venue information")
		}
		venue := strings.Split(venueInfo, "Live at ")
		appendVenue(venue[1], &performers)

		// Get set time info
		setTimeInfo, err := getEventDetails(eventDetails, -1)
		if err != nil {
			fmt.Println("Error getting set time information")
		}
		eventTime := strings.Split(setTimeInfo, "sets start at ")
		appendEventTime(eventTime[1], &performers)

		// Get set date info
		e.ForEach("div.sub-info__date-time", func(_ int, elem *colly.HTMLElement) {
			info := elem.ChildText("div.title5:first-child")
			appendEventDate(info, &performers)
		})

		// Get performers info
		e.ForEach("div.title5", func(_ int, elem *colly.HTMLElement) {
			text := elem.Text
			parts := strings.Split(text, " / ")
			if len(parts) == 2 {
				performer := strings.TrimSpace(parts[0])
				instrument := strings.TrimSpace(parts[1])
				performers[performer] = instrument
			}
		})

		// Add current time stamp
		utils.AppendCurrentTime(&performers)

		// POST data to server
		utils.PostVenueData(venue[1], &performers)
	})

	c.Visit("https://www.smallslive.com/")
}

func appendEventTitle(eventTitle string, postable *map[string]string) {
	(*postable)["eventTitle"] = eventTitle
}

func appendEventTime(setTime string, postable *map[string]string) {
	(*postable)["eventTime"] = setTime
}

func appendEventDate(eventDate string, postable *map[string]string) {
	(*postable)["eventDate"] = eventDate
}

func appendVenue(venue string, postable *map[string]string) {
	venueName := strings.ToLower(venue)
	(*postable)["venue"] = venueName
}

func getEventDetails(details []string, target int) (string, error) {
	if target == -1 {
		return details[len(details)-1], nil
	}

	if target == -2 {
		return details[len(details)-2], nil
	}

	return "", errors.New("Cannot get index")
}
