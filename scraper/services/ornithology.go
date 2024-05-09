package scraper

import (
	"github.com/gocolly/colly"
	"strings"
	"whoshittin/scraper/utils"
	"whoshittin/scraper/venueNames"
)

type OrnithologyData struct {
	EventInfo
}

const visitUrl = "https://www.ornithologyjazzclub.com/events-2"

func Ornithology(c *colly.Collector) {
	instrumentMap := map[string]string{
		"b":         "bass",
		"p":         "piano",
		"d":         "drums",
		"g":         "guitar",
		"as":        "alto sax",
		"ts":        "tenor sax",
		"bs":        "bari sax",
		"tp":        "trumpet",
		"trombone":  "trombone",
		"saxophone": "saxophone",
		"organ":     "organ",
	}
	c.OnHTML("div.eventlist-column-info", func(e *colly.HTMLElement) {
		var eventData OrnithologyData
		eventData.AppendVenue(venueNames.OrnithologyVenueName)
		linkUrl := visitUrl
		extractedHref := e.ChildAttr("a.eventlist-button", "href")
		modifiedURL := strings.Replace(linkUrl, "/events-2", "", 1)
		finalURL := modifiedURL + extractedHref
		eventData.AppendEventLink(finalURL)
		eventData.AppendCurrentTime()
		eventTitle := e.ChildText("h1 > a")
		eventData.AppendEventTitle(eventTitle)
		appendEventTime(eventTitle, &eventData)
		eventData.AppendEventImage("https://images.squarespace-cdn.com/content/v1/611849a90dfaab4317cd4c6b/86dce9c5-3273-49cb-a46d-8709b98bac56/68571516_padded_logo.png?format=1500w")

		var isElementFound bool = false
		e.ForEach("time.event-date", func(_ int, child *colly.HTMLElement) {
			if !isElementFound {
				eventData.AppendEventDate(child.Text)
				isElementFound = true
			}
		})

		e.ForEach("div.sqs-html-content > p", func(_ int, p *colly.HTMLElement) {
			var performer Performer
			originalText := p.Text
			formattedText := strings.ReplaceAll(strings.ReplaceAll(originalText, "(", " - "), ")", "")
			splitText := strings.Split(formattedText, " - ")
			performer.Name = splitText[0]
			value, ok := instrumentMap[splitText[1]]
			if ok {
				performer.Instrument = value
			} else {
				performer.Instrument = splitText[1]
			}
			eventData.AddBandMember(performer)
		})
		utils.PostVenueData(venueNames.OrnithologyVenueName, eventData)
	})
	c.Visit(visitUrl)
}

func appendEventTime(title string, eventData *OrnithologyData) {
	if title == "Jazz Dialogue Open Jam" {
		eventData.AppendEventTime("9:00PM - 12:00AM")
	} else {
		eventData.AppendEventTime("6:30PM - 8:30PM")
	}
}
