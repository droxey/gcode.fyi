package scraper

import (
	"strings"
	"time"

	"github.com/droxey/gcode.fyi/backend/utils"

	"github.com/gocolly/colly"
	"github.com/jinzhu/gorm"
)

const titleSeparator = ":"
const minLength = 3

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
		codeDescription := strings.Split(title, titleSeparator)[0]
		codeDescription = strings.Split(codeDescription, " ")
		code := codeDescription[0]
		desc := codeDescription[1]

		isValidCommand := (len(code) >= minLength) && (strings.HasPrefix(code, "M") || strings.HasPrefix(code, "G"))
		if !isValidCommand {
			return
		}
		utils.Debug(code)
		// e.ForEach("#mw-content-text > div > h4", func(_ int, el *colly.HTMLElement) {

		// 	if !strings.HasPrefix(cmdTitle, "G") || !strings.HasPrefix(cmdTitle, "M") {
		// 		return
		// 	}

		// 	// cmd := &models.Command{}
		// 	// cmd.Firmware = "Marlin"
		// 	// cmd.Version = "1.1.9-bugfix"
		// 	// cmd.SourceURL = e.Request.URL.String()
		// 	// db.FirstOrCreate(&cmd)
		// 	// commandCount++
		// })
	})

	// Handle any errors that occur.
	c.OnError(func(r *colly.Response, err error) {
		utils.CheckError(err)
	})

	c.Visit("https://reprap.org/wiki/G-code")

	return commandCount
}
