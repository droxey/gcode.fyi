package scraper

import (
	"fmt"
	"strings"
	"time"

	"github.com/droxey/gcode.fyi/backend/models"
	"github.com/droxey/gcode.fyi/backend/utils"

	"github.com/gocolly/colly"
	"github.com/jinzhu/gorm"
	cmap "github.com/orcaman/concurrent-map"
	// concurrent-map provides a high-performance solution to map's native lack of concurrent operations.
)

// RunGCODEScraper TODO
func RunGCODEScraper(db *gorm.DB) int {
	commandMap := cmap.New()

	c := colly.NewCollector(
		colly.UserAgent("gcode.fyi"),
		colly.CacheDir("./.cache"),
		colly.MaxDepth(2),
	)

	c.Limit(&colly.LimitRule{DomainGlob: "*", RandomDelay: 5 * time.Second})

	firmwares := []models.Firmware{}
	db.Where(&models.Firmware{}).Find(&firmwares)
	firmwareMap := cmap.New()
	if len(firmwares) == 0 {
		c2 := c.Clone()

		// Locate the div immediately following the title.
		// Scrape the table's contents for firmware information.
		// Select only the first 18 to prevent redundancy.
		c2.OnHTML("#mw-content-text > div > div:nth-child(16) > table > tbody > tr > th:nth-child(-n+18)", func(e *colly.HTMLElement) {
			name := e.DOM.Find("a").Text()
			if len(name) <= 3 {
				return
			}

			fw := models.Firmware{
				Name:          e.DOM.Find("a").Text(),
				RepRapWikiURL: "https://reprap.org" + e.DOM.Find("a").AttrOr("href", "/wiki"),
			}

			db.Create(&fw)
			firmwareMap.Set(fw.Name, fw)
		})

		c2.Visit("https://reprap.org/wiki/G-code")
		utils.Debug(fmt.Sprintf("=> Scraping Firmware: %d Found.", firmwareMap.Count()))
	}

	c.OnHTML("#mw-content-text > div > h4", func(e *colly.HTMLElement) {
		title := e.ChildText("span")

		// Title text contains the GCODE command and a short description
		// of the command, separated by a colon (`:`).
		// May contain multiple commands separated by an ampersand (`&`).
		commands := utils.FindGCODEInString(title)

		// Iterate over all found commands.
		for _, code := range commands {
			// Cleanup command name.
			name := strings.TrimSpace(strings.Replace(strings.Replace(strings.Replace(title, code, "", 1), ": ", "", 1), " & ", "", 1))
			if strings.HasPrefix(name, code) {
				name = strings.Replace(name, code, "", 1)
			}

			// Create a Command instance to store the data.
			cmd := &models.Command{
				Code:      code,
				Name:      name,
				SourceURL: e.Request.URL.String(),
			}

			db.Create(&cmd)
			commandMap.Set(cmd.Code, cmd)

			// firmwareTableFeatures := e.DOM.Parent().Find("h4 ~ div > table > tbody > tr > td")
			// for tdNode := range firmwareTableFeatures.Nodes {
			// 	td := firmwareTableHeaders.Eq(tdNode)
			// 	isSupported := strings.Replace(td.Text(), "???", "Unknown", 1)

			// 	foundOpts := []string{"Yes", "Mo", "Unknown"}
			// 	_, found := utils.Find(foundOpts, isSupported)
			// 	if !found {
			// 		// fw.Notes = isSupported
			// 		isSupported = "Unknown"
			// 	}

			// 	utils.Debug(isSupported)
			// }

			detailsList := e.DOM.Parent().Find("h4 ~ dl")
			for dlNode := range detailsList.Nodes {
				dl := detailsList.Eq(dlNode)
				headerText := dl.Find("dd").Text()
				if strings.Contains(headerText, "Main article:") ||
					strings.Contains(headerText, "Reserved") ||
					len(headerText) < 3 {
					return
				}

				switch headerText {
				case "Parameters":
					utils.Debug("  Parameters:")
					p := dl.Parent().Find("code").Text()
					param := models.Parameter{
						Command:     *cmd,
						Parameter:   strings.TrimSpace(p),
						Description: strings.TrimSpace(strings.Replace(headerText, p, "", 1)),
					}
					db.Create(&param)
					utils.Debug("    - Parameter: " + param.Parameter)
				case "Example":
					utils.Debug("  Example: ")
					exampleCode := e.DOM.Parent().Find("pre").Text()
					utils.Debug("    " + exampleCode)
				}
			}

		}
	})

	// Handle any errors that occur.
	c.OnError(func(r *colly.Response, err error) {
		utils.CheckError(err)
	})

	c.Visit("https://reprap.org/wiki/G-code")

	return commandMap.Count()
}
