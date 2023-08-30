package scraper

import "github.com/gocolly/colly"

func Scraper() {
	smallsLive := colly.NewCollector()
	SmallsLiveScraper(smallsLive)
}
