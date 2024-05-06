package scraper

import (
	"github.com/gocolly/colly"
	"regexp"
	"whoshittin/scraper/utils"
)

type SmokeData struct {
	EventInfo
}

func Smoke(c *colly.Collector) {
	c.OnHTML("div.details.border-b", func(e *colly.HTMLElement) {
		var eventData SmokeData
		eventData.AppendEventTime(utils.GetCurrentTime())
		eventData.AppendEventTitle(e.ChildText("h3.text-3xl"))
		eventData.AppendEventDate(e.ChildText("h4.day-of-week"))
		eventData.AppendEventImage(e.ChildAttr("img", "src"))
		eventData.AppendVenue("smoke")

		text := e.ChildText("span")
		re := regexp.MustCompile(`([^\s–]+ [^\s–]+) – ([^\n]+)`)
		matches := re.FindAllStringSubmatch(text, -1)
		for _, match := range matches {
			var performer Performer
			name := match[1]
			instrument := match[2]
			performer.Name = name
			performer.Instrument = instrument
			eventData.AddBandMember(performer)
		}

		// POST data to server
		utils.PostVenueData("smoke", eventData)
	})
	c.Visit("https://tickets.smokejazz.com/")

}
