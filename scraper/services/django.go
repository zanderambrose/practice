package scraper

import (
	"github.com/gocolly/colly"
	"strings"
	"whoshittin/scraper/utils"
	"whoshittin/scraper/venueNames"
)

type DjangoData struct {
	EventInfo
}

func Django(c *colly.Collector) {
	c.OnHTML("article.event_card", func(e *colly.HTMLElement) {
		var earlyEventData DjangoData
		var lateEventData DjangoData
		earlyEventData.AppendVenue(venueNames.Django)
		lateEventData.AppendVenue(venueNames.Django)
		earlyEventData.AppendCurrentTime()
		lateEventData.AppendCurrentTime()
		parts := strings.Split(e.ChildText("p.event__info"), "\n")
		earlyEventData.AppendEventDate(parts[0])
		earlyEventData.AppendEventTime(parts[2])
		lateEventData.AppendEventDate(parts[0])
		lateEventData.AppendEventTime(parts[2])
		earlyEventData.AppendEventTitle(e.ChildText("h3"))
		earlyEventData.AppendEventLink(e.ChildAttr("a.details-container", "href"))
		earlyEventData.AppendEventImage(e.ChildAttr("img", "src"))
		utils.PostVenueData(venueNames.Django, earlyEventData)
		lateEventData.AppendEventTitle(e.ChildText("h3"))
		lateEventData.AppendEventLink(e.ChildAttr("a.details-container", "href"))
		lateEventData.AppendEventImage(e.ChildAttr("img", "src"))
		utils.PostVenueData(venueNames.Django, lateEventData)
	})
	c.Visit("https://www.thedjangonyc.com/events")
}
