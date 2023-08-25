package scraper

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

func Scraper() {
	c := colly.NewCollector()

	allPerformers := make([]map[string]string, 0)

	c.OnHTML("div.artists", func(e *colly.HTMLElement) {
		performers := make(map[string]string)

		e.ForEach("a.title5", func(_ int, elem *colly.HTMLElement) {
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
		if i%2 == 0 {
			fmt.Println("Smalls")
		} else {
			fmt.Println("Mezzrow")
		}
		for performer, instrument := range performers {
			fmt.Printf("%s: %s;\n", performer, instrument)
		}
		fmt.Println()
	}

}
