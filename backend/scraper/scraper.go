package scraper

import (
	"strings"
	"time"

	"github.com/droxey/gcode.fyi/backend/models"
	"github.com/droxey/gcode.fyi/backend/utils"

	"github.com/gocolly/colly"
	"github.com/jinzhu/gorm"
)

// RunGCODEScraper TODO
func RunGCODEScraper(db *gorm.DB) int {
	commandCount := 0

	c := colly.NewCollector(
		colly.UserAgent("gcode.fyi"),
		colly.CacheDir("./.cache"),
		colly.MaxDepth(2),
	)

	c.Limit(&colly.LimitRule{DomainGlob: "*", RandomDelay: 5 * time.Second})

	// Scrape marlin.org for links to docs for each command.
	c.OnHTML("div.item > h2 > a", func(e *colly.HTMLElement) {
		marlinDocsLink := e.Attr("href")
		e.Request.Visit(marlinDocsLink)
	})

	c.OnHTML("#mw-content-text > div > h4", func(e *colly.HTMLElement) {
		title := e.DOM.Find("span").Text()

		// Title text contains the GCODE command and a short description
		// of the command, separated by a colon (`:`).
		// May contain multiple commands separated by an ampersand (`&`).
		commands := utils.FindGCODEInString(title)

		// Iterate over all found commands.
		for _, code := range commands {
			// Create a Command instance to store the data.
			commandModel := &models.Command{}
			commandModel.Code = code
			commandModel.SourceURL = e.Request.URL.String()
			commandModel.ShortDescription = strings.TrimSpace(strings.Replace(strings.Replace(strings.Replace(title, code, "", 1), ": ", "", 1), " & ", "", 1))
			commandCount = commandCount + 1
			utils.Debug(commandModel.Code + ": " + commandModel.ShortDescription)
		}
	})

	// Handle any errors that occur.
	c.OnError(func(r *colly.Response, err error) {
		utils.CheckError(err)
	})

	// Start scraping.
	c.Visit("https://reprap.org/wiki/G-code")

	return commandCount
}
