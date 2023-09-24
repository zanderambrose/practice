package scraper

import (
	"fmt"
	"github.com/gocolly/colly"
	"practice/pkg/utils"
	"strings"
	"unicode/utf8"
)

func Vanguard(c *colly.Collector) {
	data := make(map[string]string, 0)
	c.OnHTML("div.event-details", func(e *colly.HTMLElement) {
		if e.Index == 0 {

			// Add band name
			bandName := e.ChildText("h2")
			data["bandName"] = bandName

			// Add current time
			utils.AppendCurrentTime(&data)

			// Add band members
			e.ForEach("h4", func(_ int, bandMember *colly.HTMLElement) {
				text := bandMember.Text
				dash := '–'
				parts := strings.IndexRune(text, dash)
				if parts != -1 {
					performer := strings.TrimSpace(text[:parts])
					instrument := strings.TrimSpace(text[parts+utf8.RuneLen(dash):])
					data[performer] = instrument
				}
			})
		}

	})

	c.Visit("https://villagevanguard.com/")
	fmt.Println(data)
}
