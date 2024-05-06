package scraper

import "github.com/gocolly/colly"

type ScraperFunc func(*colly.Collector)

var ScraperMap = map[string]ScraperFunc{
	"smalls":   SmallsLiveScraper,
	"vangaurd": Vanguard,
	"django":   Django,
	"smoke":    Smoke,
}

func Scraper() {
	c := colly.NewCollector()
	for key := range ScraperMap {
		scraperFunc := ScraperMap[key]
		scraperFunc(c)
	}
}
