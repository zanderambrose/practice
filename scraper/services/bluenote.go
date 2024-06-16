package scraper

import (
	"github.com/gocolly/colly"
	"sync"
	"whoshittin/scraper/utils"
	"whoshittin/scraper/venueNames"
)

func BlueNote(c *colly.Collector, w *sync.WaitGroup) {
	c.OnHTML("div.inner", func(e *colly.HTMLElement) {
		dayOfTheMonth := e.ChildText("div.day")
		e.ForEach("div.day-wrap", func(_ int, item *colly.HTMLElement) {
			var eventData EventInfo
			detailsUrl := item.ChildAttr("a", "href")
			eventData.AppendVenue(venueNames.BlueNote)
			eventData.AppendCurrentTime()
			eventData.AppendEventTitle(item.ChildText("h3"))
			eventData.AppendEventLink(detailsUrl)
			eventData.AppendEventDate(dayOfTheMonth)

			var allShowTimes string
			item.ForEach("time", func(_ int, time *colly.HTMLElement) {
				eventTime := time.Text
				if len(allShowTimes) != 0 {
					allShowTimes += " & "
				}
				allShowTimes += eventTime
			})
			eventData.AppendEventTime(allShowTimes)
			findEventImage(&eventData, detailsUrl)
		})
	})
	c.Visit("https://www.bluenotejazz.com/nyc/shows/?calendar_view")
	defer w.Done()
}

func findEventImage(eventData *EventInfo, visitUrl string) {
	c := colly.NewCollector()
	c.OnHTML("main", func(e *colly.HTMLElement) {
		image := e.ChildAttr("img.the-group-image", "src")
		eventData.AppendEventImage(image)
		utils.PostVenueData(venueNames.BlueNote, eventData)
	})
	c.Visit(visitUrl)
}
