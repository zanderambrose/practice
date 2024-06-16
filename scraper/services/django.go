package scraper

import (
	"github.com/gocolly/colly"
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
		times := e.ChildText("p.event_card__time-pair")
		eventTimes := strings.Replace(times, "|", "&", 1)
		eventData.AppendEventTime(eventTimes)
		utils.PostVenueData(venueNames.Django, &eventData)
	})
	c.Visit("https://www.thedjangonyc.com/events")
	defer w.Done()
}
