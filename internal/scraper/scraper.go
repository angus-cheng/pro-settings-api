package scraper

import (
	"github.com/gocolly/colly"
)

func Run() {
	c := colly.NewCollector()
	c.OnHTML("table tr", func(e *colly.HTMLElement) {
		ps := parsePlayer(e)
		db.savePlayer(ps)
	})
	c.Visit("https://prosettings.net/list/valorant/")
}