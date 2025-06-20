package scraper

import (
	"github.com/angus-cheng/pro-settings-api/internal/db"
	"github.com/gocolly/colly"
)

func Run() {
	c := colly.NewCollector()
	c.OnHTML("table tr", func(e *colly.HTMLElement) {
		ps := ParsePlayer(e)
		db.SavePlayer(ps)
	})
	c.Visit("https://prosettings.net/list/valorant/")
}