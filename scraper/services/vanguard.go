package scraper

import (
	"github.com/gocolly/colly"
	"regexp"
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
			trimmedEventText := strings.TrimSpace(replaceWhitespace(e.ChildText("div.event-details > h3")))
			eventData.AppendEventDate(trimmedEventText)
			eventData.AppendEventImage(e.ChildAttr("img", "src"))
			eventData.AppendVenue(vanguardVenueName)

			// TODO - HANDLE DIFFERENT BANDS FOR DIFFERENT DATES IN H4 LOOPS
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
					// TODO - Handle Vanguard jazz orchestra band member formatting
				}
			})
			utils.PostVenueData(vanguardVenueName, eventData)
		}
	})

	c.Visit("https://villagevanguard.com/")
}

func replaceWhitespace(input string) string {
	// replace whitespace characters with a single space
	re := regexp.MustCompile(`\s+`)
	return re.ReplaceAllString(input, " ")
}
