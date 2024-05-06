package scraper

import (
	"fmt"
	"github.com/gocolly/colly"
	"whoshittin/scraper/utils"
)

func Django(c *colly.Collector) {
	earlySet := make(map[string]string, 0)
	lateSet := make(map[string]string, 0)

	c.OnHTML("article", func(e *colly.HTMLElement) {
		if e.Index == 1 {
			// Add band name
			AppendEventTitle(&earlySet, e.ChildText("h3"))

			// Add details url
			appendDetailsUrl(&earlySet, e.ChildAttr("a.details-container", "href"))
		}

		if e.Index == 2 {
			// Add band name
			AppendEventTitle(&lateSet, e.ChildText("h3"))

			// Add details url
			appendDetailsUrl(&lateSet, e.ChildAttr("a.details-container", "href"))
		}
	})

	c.Visit("https://www.thedjangonyc.com/events")
	fmt.Println("earlySet: ", earlySet)
	fmt.Println("lateSet: ", lateSet)
	utils.PostVenueData("django", &earlySet)
	utils.PostVenueData("django", &lateSet)
}

func AppendEventTitle(setData *map[string]string, bandName string) {
	(*setData)["eventTitle"] = bandName
}

func appendDetailsUrl(setData *map[string]string, url string) {
	(*setData)["detailsUrl"] = url
}
