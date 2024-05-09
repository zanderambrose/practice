package scraper

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
	"time"
	"whoshittin/scraper/utils"
)

type EventInfo struct {
	Venue       string      `json:"venue" bson:"venue"`
	Band        []Performer `json:"band" bson:"band"`
	EventTime   string      `json:"eventTime" bson:"eventTime"`
	EventDate   EventDate   `json:"eventDate" bson:"eventDate"`
	CurrentTime string      `json:"currentTime" bson:"currentTime"`
	EventTitle  string      `json:"eventTitle" bson:"eventTitle"`
	EventImage  string      `json:"eventImage" bson:"eventImage"`
	EventLink   string      `json:"eventLink" bson:"eventLink"`
}

type Performer struct {
	Instrument string `json:"instrument"`
	Name       string `json:"name"`
}

type EventDate struct {
	EventDate           string    `json:"eventDate" bson:"eventDate"`
	EventDateNormalized time.Time `json:"eventDateNormalized" bson:"eventDateNormalized"`
}

func (data *EventInfo) AppendEventTitle(eventTitle string) {
	data.EventTitle = eventTitle
}

func (data *EventInfo) AppendEventLink(eventLink string) {
	data.EventLink = eventLink
}

func (data *EventInfo) AppendEventTime(setTime string) {
	data.EventTime = setTime
}

func (data *EventInfo) AppendEventDate(eventDate string) {
	data.EventDate.EventDate = eventDate
	normalizedDate, err := utils.NormalizeDate(eventDate, data.Venue)
	if err != nil {
		fmt.Println("Normalized Date error", err)
	}
	data.EventDate.EventDateNormalized = normalizedDate
}

func (data *EventInfo) AppendVenue(venue string) {
	data.Venue = strings.ToLower(venue)
}

func (data *EventInfo) AddBandMember(performer Performer) {
	data.Band = append(data.Band, performer)
}

func (data *EventInfo) AppendEventImage(imgSrc string) {
	if strings.Contains(imgSrc, "https") {
		data.EventImage = imgSrc
	} else {
		data.EventImage = "https:" + imgSrc
	}
}

func (data *EventInfo) AppendCurrentTime() {
	time := utils.GetCurrentTime()
	data.CurrentTime = time
}

type ScraperFunc func(*colly.Collector)

var ScraperMap = map[string]ScraperFunc{
	"smalls":      SmallsLiveScraper,
	"mezzrow":     SmallsLiveScraper,
	"vanguard":    Vanguard,
	"django":      Django,
	"smoke":       Smoke,
	"ornithology": Ornithology,
}

func Scraper() {
	c := colly.NewCollector()
	for key := range ScraperMap {
		scraperFunc := ScraperMap[key]
		scraperFunc(c)
	}
}
