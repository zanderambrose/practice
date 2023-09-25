package scraper

import (
	"fmt"
	"practice/pkg/utils"

	"github.com/gocolly/colly"
	// "strings"
	// "unicode/utf8"
)

func Django(c *colly.Collector) {
	earlySet := make(map[string]string, 0)
	lateSet := make(map[string]string, 0)

	c.OnHTML("article", func(e *colly.HTMLElement) {
		if e.Index == 1 {
			// Add band name
			appendBandName(&earlySet, e.ChildText("h3"))

			// Add details url
			appendDetailsUrl(&earlySet, e.ChildAttr("a.details-container", "href"))

			// Add current time
			utils.AppendCurrentTime(&earlySet)
		}

		if e.Index == 2 {
			// Add band name
			appendBandName(&lateSet, e.ChildText("h3"))

			// Add details url
			appendDetailsUrl(&lateSet, e.ChildAttr("a.details-container", "href"))

			// Add current time
			utils.AppendCurrentTime(&lateSet)
		}
	})

	c.Visit("https://www.thedjangonyc.com/events")
	fmt.Println("earlySet: ", earlySet)
	fmt.Println("lateSet: ", lateSet)
	// utils.PostVenueData("django", &earlySet)
	// utils.PostVenueData("django", &lateSet)
}

func appendBandName(setData *map[string]string, bandName string) {
	(*setData)["bandName"] = bandName
}

func appendDetailsUrl(setData *map[string]string, url string) {
	(*setData)["detailsUrl"] = url
}
