package main

import (
	"os"
	"time"

	"github.com/droxey/gcode.fyi/backend/db"
	"github.com/droxey/gcode.fyi/backend/scraper"
	"github.com/droxey/gcode.fyi/backend/utils"
)

func main() {
	d := db.New()
	db.AutoMigrate(d)

	start := time.Now()
	commandCount := scraper.RunGCODEScraper(d)
	utils.ShowSummary(start, commandCount)
	os.Exit(1)
}
