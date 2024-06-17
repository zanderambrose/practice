package scraper

import (
	"github.com/gocolly/colly"
	"regexp"
	"strings"
	"sync"
	"whoshittin/scraper/utils"
	"whoshittin/scraper/venueNames"
)

func Django(c *colly.Collector, w *sync.WaitGroup) {
	c.OnHTML("article.event_card", func(e *colly.HTMLElement) {
		var eventData EventInfo
		eventData.AppendVenue(venueNames.Django)
		eventData.AppendCurrentTime()
		eventDate := e.ChildText("p.event__date")
		eventData.AppendEventDate(eventDate)
		eventData.AppendEventTitle(e.ChildText("h3"))
		eventData.AppendEventLink(e.ChildAttr("a.details-container", "href"))
		eventData.AppendEventImage(e.ChildAttr("img", "src"))
		times := e.ChildText("p.event_card__time-pair > span")
		strippedTimes := strings.ReplaceAll(times, "|", "")
		pattern := `(?i)(\d{1,2}:\d{2}[ap]m)`
		re, err := regexp.Compile(pattern)
		if err != nil {
			println("Error compiling regexp: ", err)
		}
		splitTimes := re.FindAllString(strippedTimes, -1)

		var allTimes string

		for i, v := range splitTimes {
			if i < len(splitTimes)-1 {
				allTimes += v + "&"
			} else {
				allTimes += v
			}
		}
		eventData.AppendEventTime(allTimes)
		utils.PostVenueData(venueNames.Django, &eventData)
	})
	c.Visit("https://www.thedjangonyc.com/events")
	defer w.Done()
}
