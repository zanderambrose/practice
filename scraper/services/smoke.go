package scraper

import (
	"github.com/gocolly/colly"
	"regexp"
	"strings"
	"whoshittin/scraper/utils"
)

type SmokeData struct {
	EventInfo
}

const smokeVenueName = "smoke"

func Smoke(c *colly.Collector) {
	c.OnHTML("div.details.border-b", func(e *colly.HTMLElement) {
		var eventData SmokeData
		eventData.AppendCurrentTime()
		eventData.AppendEventTitle(e.ChildText("h3.text-3xl"))
		eventData.AppendEventDate(e.ChildText("h4.day-of-week"))
		eventData.AppendEventImage(e.ChildAttr("img", "src"))
		eventData.AppendVenue(smokeVenueName)

		descriptionText := e.ChildText("span")
		re := regexp.MustCompile(`([^\s–]+ [^\s–]+) – ([^\n]+)`)
		matches := re.FindAllStringSubmatch(descriptionText, -1)
		for _, match := range matches {
			var performer Performer
			name := match[1]
			instrument := match[2]
			performer.Name = name
			performer.Instrument = instrument
			eventData.AddBandMember(performer)
		}

		var allShowTimes string
		e.ForEach("button", func(_ int, time *colly.HTMLElement) {
			eventTime := time.Text
			showIndex := strings.Index(eventTime, " SHOW")

			showTime := eventTime[:showIndex]

			if len(allShowTimes) != 0 {
				allShowTimes += " & "
			}
			allShowTimes += showTime
		})
		eventData.AppendEventTime(allShowTimes)
		utils.PostVenueData(smokeVenueName, eventData)
	})
	c.Visit("https://tickets.smokejazz.com/")

}
