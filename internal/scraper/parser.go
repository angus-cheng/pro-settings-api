package scraper

import (
	"github.com/gocolly/colly"
)

func parsePlayer(e *colly.HTMLElement) models.Player {
	return Models.Player{
		Name:        e.ChildText("td.name"),
		Sensitivity: parseFloat(e.ChildText("td.sens")),
	}
}