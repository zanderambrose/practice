package scraper

import (
	"github.com/gocolly/colly"
	// "strings"
	"whoshittin/scraper/utils"
	"whoshittin/scraper/venueNames"
)

func Zinc(c *colly.Collector) {
	c.OnHTML("body", func(e *colly.HTMLElement) {
		e.ForEach("div.edgtf-row-grid-section-wrapper", func(idx int, container *colly.HTMLElement) {
			if idx == 0 {
				return
			}
			container.ForEach("div.edgtf-el-item", func(idx int, item *colly.HTMLElement) {
				// FOR DEVELOPMENT ONLY
				if idx > 0 {
					return
				}
				var eventData EventInfo
				eventData.AppendVenue(venueNames.Zinc)
				eventData.AppendCurrentTime()
				eventData.AppendEventTitle(item.ChildText("h4"))
				detailsUrl := item.ChildAttr("a", "href")
				eventData.AppendEventLink(detailsUrl)

				weekday := item.ChildText("span.edgtf-el-item-weekday")
				month := item.ChildText("span.edgtf-el-item-month")
				day := item.ChildText("span.edgtf-el-item-day")
				eventData.AppendEventDate(weekday + " " + month + " " + day)
				visitEventDetails(&eventData, detailsUrl)
			})
		})
	})
	c.Visit("https://www.zincbar.com/shows/")
}

func visitEventDetails(eventData *EventInfo, visitUrl string) {
	c := colly.NewCollector()
	c.OnHTML("div.offbeat-event-top-holder", func(e *colly.HTMLElement) {
		image := e.ChildAttr("img", "src")
		eventData.AppendEventImage(image)
		// eventData.AppendEventTime(allShowTimes)
		utils.PostVenueData(venueNames.Zinc, eventData)
	})
	c.Visit(visitUrl)
}
