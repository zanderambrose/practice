package scraper

import (
	"github.com/gocolly/colly"
	"strings"
	"unicode/utf8"
	"whoshittin/scraper/utils"
)

func Vanguard(c *colly.Collector) {
	data := make(map[string]string, 0)
	c.OnHTML("div.event-details", func(e *colly.HTMLElement) {
		if e.Index == 0 {

			// Add band name
			AppendBandName(&data, e.ChildText("h2"))

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
	utils.PostVenueData("vanguard", &data)
}
