package scraper

import (
	"errors"
	"fmt"
	"github.com/gocolly/colly"
	"strings"
	"whoshittin/scraper/utils"
)

func SmallsLiveScraper(c *colly.Collector) {
	c.OnHTML("article.event-display-today-and-tomorrow", func(e *colly.HTMLElement) {
		var eventData EventInfo

		ariaLabel := e.ChildAttr("a", "aria-label")
		eventDetails := strings.Split(ariaLabel, ", ")

		// Get venue info
		venueInfo, err := getEventDetails(eventDetails, -2)
		if err != nil {
			fmt.Println("Error getting venue information")
		}
		venue := strings.Split(venueInfo, "Live at ")
		eventData.AppendVenue(venue[1])

		// Get event img
		e.ForEach("div.event-picture", func(_ int, elem *colly.HTMLElement) {
			imgSrc := elem.ChildAttr("img", "src")
			eventData.AppendEventImage(imgSrc)
		})

		// Get event title
		eventTitle := e.ChildText("p.event-info-title")
		eventData.AppendEventTitle(eventTitle)

		// Get set time info
		setTimeInfo, err := getEventDetails(eventDetails, -1)
		if err != nil {
			fmt.Println("Error getting set time information")
		}
		eventTime := strings.Split(setTimeInfo, "sets start at ")
		eventData.AppendEventTime(eventTime[1])

		// Get set date info
		e.ForEach("div.sub-info__date-time", func(_ int, elem *colly.HTMLElement) {
			eventData.AppendEventDate(elem.ChildText("div.title5:first-child"))
		})

		// Get performers info
		e.ForEach("div.title5", func(_ int, elem *colly.HTMLElement) {
			var performer Performer
			text := elem.Text
			parts := strings.Split(text, " / ")
			if len(parts) == 2 {
				name := strings.TrimSpace(parts[0])
				instrument := strings.TrimSpace(parts[1])
				performer.Name = name
				performer.Instrument = instrument
				eventData.AddBandMember(performer)
			}
		})

		eventData.AppendCurrentTime()

		// POST data to server
		utils.PostVenueData(venue[1], eventData)
	})

	c.Visit("https://www.smallslive.com/")
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
