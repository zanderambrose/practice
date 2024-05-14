package scraper

import (
	"fmt"
	"github.com/gocolly/colly"
	"regexp"
	"whoshittin/scraper/utils"
	"whoshittin/scraper/venueNames"
)

func BlueNote(c *colly.Collector) {
	c.OnHTML("div.day-wrap", func(e *colly.HTMLElement) {
		// e.ForEach("div.day-wrap", func(_ int, time *colly.HTMLElement) {

		// })
		var eventData EventInfo
		eventData.AppendVenue(venueNames.BlueNote)
		eventData.AppendCurrentTime()
		eventData.AppendEventTitle(e.ChildText("h3 > a"))
		eventData.AppendEventDate(e.ChildText("h4.day-of-week"))
		eventData.AppendEventLink(e.ChildAttr("a", "href"))

		backgroundImage := e.ChildAttr("div.the-image", "style")
		re := regexp.MustCompile(`url\("(.+?)"\)`)
		matches := re.FindStringSubmatch(backgroundImage)
		fmt.Println("matches:", matches)
		if len(matches) >= 2 {
			url := matches[1]
			// Print or use the extracted URL
			fmt.Println("URL:", url)
		}
		eventData.AppendEventImage(e.ChildAttr("div.the-image", "src"))

		var allShowTimes string
		e.ForEach("time", func(_ int, time *colly.HTMLElement) {
			eventTime := time.Text
			if len(allShowTimes) != 0 {
				allShowTimes += " & "
			}
			allShowTimes += eventTime
		})
		eventData.AppendEventTime(allShowTimes)
		utils.PostVenueData(venueNames.BlueNote, eventData)
	})
	c.Visit("https://www.bluenotejazz.com/nyc/shows/?calendar_view")

}
