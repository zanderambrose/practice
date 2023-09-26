package scraper

import (
	"fmt"
	"github.com/gocolly/colly"
	// "practice/pkg/utils"
	"strings"
)

func Smoke(c *colly.Collector) {
	data := make(map[string]string, 0)

	c.OnHTML("div", func(e *colly.HTMLElement) {
		text := strings.TrimSpace(e.Text)
		fmt.Println("on html: ", text)
	})

	err := c.Visit("https://tickets.smokejazz.com")

	if err != nil {
		fmt.Println("Error:", err)
	}
	// utils.PostVenueData("vanguard", &data)
	fmt.Println("data: ", data)
}
