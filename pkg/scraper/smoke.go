package scraper

import (
	"github.com/gocolly/colly"
	"practice/pkg/utils"
)

type SmokeData struct {
	EventTitle      string   `json:"eventTitle"`
	EventImage      string   `json:"eventImage"`
	EventTime       []string `json:"eventTime"`
	EventDate       string   `json:"eventDate"`
	CurrentTime     string   `json:"currentTime"`
	Venue           string   `json:"venue"`
	BandDescription string   `json:"bandDescription"`
}

func Smoke(c *colly.Collector) {
	c.OnHTML("div.show.border-b", func(e *colly.HTMLElement) {
		eventData := SmokeData{
			EventTitle:      e.ChildText("h3.text-3xl"),
			EventImage:      e.ChildAttr("img", "src"),
			EventDate:       e.ChildText("h4.day-of-week"),
			CurrentTime:     utils.GetCurrentTime(),
			Venue:           "smoke",
			BandDescription: e.ChildText("p.inline"),
		}

		e.ForEach("button", func(_ int, elem *colly.HTMLElement) {
			eventData.EventTime = append(eventData.EventTime, elem.Text)
		})

		// POST data to server
		utils.PostVenueData("smoke", eventData)
	})
	c.Visit("https://tickets.smokejazz.com/")

}
