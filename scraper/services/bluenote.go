package scraper

import (
	"github.com/gocolly/colly"
	// "regexp"
	"whoshittin/scraper/utils"
	"whoshittin/scraper/venueNames"
)

// TODO - Append event data and append event image

func BlueNote(c *colly.Collector) {
	c.OnHTML("div.inner", func(e *colly.HTMLElement) {
		e.ForEach("div.day-wrap", func(_ int, item *colly.HTMLElement) {
			var eventData EventInfo
			eventData.AppendVenue(venueNames.BlueNote)
			eventData.AppendCurrentTime()
			eventData.AppendEventTitle(item.ChildText("h3"))
			// eventData.AppendEventDate(e.ChildText("h4.day-of-week"))
			eventData.AppendEventLink(item.ChildAttr("a", "href"))

			// backgroundImage := item.ChildAttr("div.the-image", "style")
			// re := regexp.MustCompile(`url\("(.+?)"\)`)
			// matches := re.FindStringSubmatch(backgroundImage)
			// fmt.Println("matches:", matches)
			// if len(matches) >= 2 {
			// 	url := matches[1]
			// 	// Print or use the extracted URL
			// 	fmt.Println("URL:", url)
			// }
			// eventData.AppendEventImage(item.ChildAttr("div.the-image", "src"))

			var allShowTimes string
			item.ForEach("time", func(_ int, time *colly.HTMLElement) {
				eventTime := time.Text
				if len(allShowTimes) != 0 {
					allShowTimes += " & "
				}
				allShowTimes += eventTime
			})
			eventData.AppendEventTime(allShowTimes)
			utils.PostVenueData(venueNames.BlueNote, eventData)
		})
	})
	c.Visit("https://www.bluenotejazz.com/nyc/shows/?calendar_view")
}
