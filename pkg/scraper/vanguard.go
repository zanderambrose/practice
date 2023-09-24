package scraper

import (
	"fmt"
	"github.com/gocolly/colly"
)

func Vanguard(c *colly.Collector) {
	c.OnHTML("div.event-details", func(e *colly.HTMLElement) {
		if e.Index == 0 {
			fmt.Println("e: ", e.Text)
		}
	})

	c.Visit("https://villagevanguard.com/")
}
