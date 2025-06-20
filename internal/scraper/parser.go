package scraper

import (
	"strconv"
	"strings"

	"github.com/angus-cheng/pro-settings-api/internal/models"
	"github.com/gocolly/colly"
)

func ParsePlayer(e *colly.HTMLElement) models.Player {
	return models.Player{
		Name:        e.ChildText("td.name"),
		Sensitivity: ParseFloat(e.ChildText("td.sens")),
	}
}

func ParseFloat(text string) float64 {
	text = strings.TrimSpace(strings.ReplaceAll(text, ",", ""))
	f, err := strconv.ParseFloat(text, 64)
	if (err != nil) {
		return 0.0
	}
	return f
}