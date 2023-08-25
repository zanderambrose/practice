package scraper

import (
	"fmt"
	"github.com/gocolly/colly"
)

func Scraper() {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://www.smallslive.com/")
}
