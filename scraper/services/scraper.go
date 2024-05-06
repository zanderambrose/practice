package scraper

import (
	"github.com/gocolly/colly"
	"strings"
	"whoshittin/scraper/utils"
)

type ScraperFunc func(*colly.Collector)

type EventInfo struct {
	Venue       string      `json:"venue" bson:"venue"`
	Band        []Performer `json:"band" bson:"band"`
	EventTime   string      `json:"eventTime" bson:"eventTime"`
	EventDate   string      `json:"eventDate" bson:"eventDate"`
	CurrentTime string      `json:"currentTime" bson:"currentTime"`
	EventTitle  string      `json:"eventTitle" bson:"eventTitle"`
	EventImage  string      `json:"eventImage" bson:"eventImage"`
}

type Performer struct {
	Instrument string `json:"instrument"`
	Name       string `json:"name"`
}

func (data *EventInfo) AppendEventTitle(eventTitle string) {
	data.EventTitle = eventTitle
}

func (data *EventInfo) AppendEventTime(setTime string) {
	data.EventTime = setTime
}

func (data *EventInfo) AppendEventDate(eventDate string) {
	data.EventDate = eventDate
}

func (data *EventInfo) AppendVenue(venue string) {
	venueName := strings.ToLower(venue)
	data.Venue = venueName
}

func (data *EventInfo) AddBandMember(performer Performer) {
	data.Band = append(data.Band, performer)
}

func (data *EventInfo) AppendEventImage(imgSrc string) {
	data.EventImage = "https:" + imgSrc
}

func (data *EventInfo) AppendCurrentTime() {
	time := utils.GetCurrentTime()
	data.CurrentTime = time
}

var ScraperMap = map[string]ScraperFunc{
	"smalls":   SmallsLiveScraper,
	"vangaurd": Vanguard,
	"django":   Django,
	"smoke":    Smoke,
}

func Scraper() {
	c := colly.NewCollector()
	for key := range ScraperMap {
		scraperFunc := ScraperMap[key]
		scraperFunc(c)
	}
}
