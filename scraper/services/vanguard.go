package scraper

import (
	"github.com/gocolly/colly"
	"strings"
	"whoshittin/scraper/utils"
)

type VanguardData struct {
	EventInfo
}

const performanceTime = "8:00 PM & 10:00 PM"
const vanguardVenueName = "vanguard"
const vanguardJazzOrchestra = "VANGUARD JAZZ ORCHESTRA"

func Vanguard(c *colly.Collector) {
	c.OnHTML("div.container", func(e *colly.HTMLElement) {
		var eventData SmokeData
		eventTitle := e.ChildText("div.event-details > h2")
		if eventTitle != "COMING SOON!" && eventTitle != "" {
			eventData.AppendCurrentTime()
			eventData.AppendEventTitle(eventTitle)
			eventData.AppendEventTime(performanceTime)
			eventData.AppendEventDate(e.ChildText("div.event-details > h3"))
			eventData.AppendEventImage(e.ChildAttr("img", "src"))
			eventData.AppendVenue(vanguardVenueName)

			e.ForEach("h4", func(_ int, bandMember *colly.HTMLElement) {
				var performer Performer
				if eventTitle != vanguardJazzOrchestra {
					text := bandMember.Text
					parts := strings.Split(text, "–")
					if len(parts) == 2 {
						name := strings.TrimSpace(parts[0])
						instrument := strings.TrimSpace(parts[1])
						performer.Name = name
						performer.Instrument = instrument
						eventData.AddBandMember(performer)
					}
				} else {

				}
			})
			utils.PostVenueData(vanguardVenueName, eventData)
		}
	})

	c.Visit("https://villagevanguard.com/")
}
