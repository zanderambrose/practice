package scraper

import (
	"fmt"
	"github.com/gocolly/colly"
	"regexp"
	"strings"
	"sync"
	"time"
	"whoshittin/scraper/utils"
	"whoshittin/scraper/venueNames"
)

type VanguardEventInfo struct {
	EventInfo
}

func (vanguardEvent *VanguardEventInfo) PostVanguardData(eventDate string) {
	if strings.Contains(strings.ToLower(eventDate), "every monday night") {
		vanguardEvent.EventDate.FormattedDate = eventDate
		vanguardEvent.EventDate.Date = time.Now().Local().Truncate(24 * time.Hour)
		utils.PostVenueData(venueNames.Vanguard, vanguardEvent)
	}
	eventDates := buildDateSlice(eventDate)

	for i := 0; i < len(eventDates); i++ {
		vanguardEvent.EventDate.Date = eventDates[i]
		vanguardEvent.EventDate.FormattedDate = eventDates[i].Format(utils.STANDARD_DATE_REPRESENTATION_LAYOUT)
		utils.PostVenueData(venueNames.Vanguard, vanguardEvent)
	}
}

const performanceTime = "8:00 PM & 10:00 PM"
const vanguardJazzOrchestra = "VANGUARD JAZZ ORCHESTRA"

func Vanguard(c *colly.Collector, w *sync.WaitGroup) {
	c.OnHTML("div.container", func(e *colly.HTMLElement) {
		var eventData VanguardEventInfo
		eventTitle := e.ChildText("div.event-details > h2")
		if eventTitle != "COMING SOON!" && eventTitle != "" {
			eventData.AppendVenue(venueNames.Vanguard)
			eventData.AppendCurrentTime()
			eventData.AppendEventTitle(eventTitle)
			eventData.AppendEventTime(performanceTime)
			eventData.AppendEventImage(e.ChildAttr("img", "src"))
			eventData.AppendEventLink(e.ChildAttr("a.btn-primary", "href"))

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
			trimmedEventText := strings.TrimSpace(replaceWhitespace(e.ChildText("div.event-details > h3")))
			eventData.PostVanguardData(trimmedEventText)
		}
	})

	c.Visit("https://villagevanguard.com/")
	defer w.Done()
}

func replaceWhitespace(input string) string {
	// replace whitespace characters with a single space
	re := regexp.MustCompile(`\s+`)
	return re.ReplaceAllString(input, " ")
}

func buildDateSlice(dateString string) []time.Time {
	parts := strings.Split(dateString, " - ")
	if len(parts) != 2 {
		fmt.Println("Invalid date range format: ", parts)
		return []time.Time{}
	}

	startStr := parts[0]
	endStr := parts[1]

	startMonth, startDay := utils.ParseDate(startStr)
	endMonth, endDay := utils.ParseDate(endStr)

	start := time.Date(time.Now().Year(), startMonth, startDay, 0, 0, 0, 0, time.UTC)
	end := time.Date(time.Now().Year(), endMonth, endDay, 0, 0, 0, 0, time.UTC)

	var dates []time.Time

	for d := start; d.Before(end.AddDate(0, 0, 1)); d = d.AddDate(0, 0, 1) {
		dates = append(dates, d)
	}
	return dates
}
