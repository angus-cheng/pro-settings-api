package main

import (
	"log"

	"github.com/angus-cheng/pro-settings-api/internal/db"
	"github.com/angus-cheng/pro-settings-api/internal/scraper"
)

func main() {
	db.Init()
	if db.DB == nil {
		log.Fatal("Database not initialized")
	}
	scraper.Run()
}