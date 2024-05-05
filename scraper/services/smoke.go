package scraper

import (
	"github.com/gocolly/colly"
	"strings"
	"whoshittin/scraper/utils"
)

type SmokeData struct {
	EventTitle      string   `json:"eventTitle" bson:"eventTitle"`
	EventImage      string   `json:"eventImage" bson:"eventImage"`
	EventTime       []string `json:"eventTime" bson:"eventTime"`
	EventDate       string   `json:"eventDate" bson:"eventDate"`
	CurrentTime     string   `json:"currentTime" bson:"currentTime"`
	Venue           string   `json:"venue" bson:"venue"`
	BandDescription string   `json:"bandDescription" bson:"bandDescription"`
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
			eventData.EventTime = append(eventData.EventTime, strings.TrimSpace(elem.Text))
		})

		// POST data to server
		utils.PostVenueData("smoke", eventData)
	})
	c.Visit("https://tickets.smokejazz.com/")

}
