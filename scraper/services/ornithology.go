package scraper

import (
	"github.com/gocolly/colly"
	"strings"
	"whoshittin/scraper/utils"
)

type OrnithologyData struct {
	EventInfo
}

const ornithologyVenueName = "ornithology"
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
		linkUrl := visitUrl
		extractedHref := e.ChildAttr("a.eventlist-button", "href")
		modifiedURL := strings.Replace(linkUrl, "/events-2", "", 1)
		finalURL := modifiedURL + extractedHref
		eventData.AppendEventLink(finalURL)

		eventData.AppendCurrentTime()
		eventData.AppendVenue(ornithologyVenueName)
		eventData.AppendEventTitle(e.ChildText("h1 > a"))
		eventData.AppendEventDate(e.ChildText("time.event-date"))
		eventData.AppendEventImage(e.ChildAttr("img", "src"))

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
		utils.PostVenueData(ornithologyVenueName, eventData)
	})
	c.Visit(visitUrl)
}
