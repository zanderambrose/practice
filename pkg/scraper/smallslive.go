package scraper

import (
	"github.com/gocolly/colly"
	"practice/pkg/utils"
	"strconv"
	"strings"
)

func SmallsLiveScraper(c *colly.Collector) {
	allPerformers := make([]map[string]string, 0)

	c.OnHTML("div.artists", func(e *colly.HTMLElement) {
		performers := make(map[string]string)

		e.ForEach("div.title5", func(_ int, elem *colly.HTMLElement) {
			text := elem.Text
			parts := strings.Split(text, " / ")
			if len(parts) == 2 {
				performer := strings.TrimSpace(parts[0])
				instrument := strings.TrimSpace(parts[1])
				performers[performer] = instrument
			}
		})

		allPerformers = append(allPerformers, performers)
	})

	c.Visit("https://www.smallslive.com/")

	for i, performers := range allPerformers {
		appendIsEarlySet(i, &performers)
		venue := appendVenue(i, &performers)
		utils.AppendCurrentTime(&performers)
		utils.PostVenueData(venue, &performers)
	}

}

func appendIsEarlySet(index int, postable *map[string]string) {
	if index > 1 {
		(*postable)["isEarlySet"] = strconv.FormatBool(false)
	} else {
		(*postable)["isEarlySet"] = strconv.FormatBool(true)
	}
}

func appendVenue(index int, postable *map[string]string) string {
	if index%2 == 0 {
		(*postable)["venue"] = "smalls"
		return "smalls"
	}

	(*postable)["venue"] = "mezzrow"
	return "mezzrow"
}
