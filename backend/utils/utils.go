package utils

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	. "github.com/logrusorgru/aurora"
)

func Debug(msg string) {
	fmt.Println(Sprintf("%s %s", White("[DEBUG]").BgGray(8-1), Blue(msg)))
}

func CheckError(err error) {
	if err != nil {
		fmt.Println(Sprintf("%s %s", White("[ERROR]").BgRed(), White(err)))
	}
}

// FindGCODEInString returns all instances of valid commands within any string.
func FindGCODEInString(str string) []string {
	var matches []string
	var re = regexp.MustCompile(`(?m)[MG]\d{1,3}`)

	for _, match := range re.FindAllString(str, -1) {
		if isValidGCODE(match) {
			matches = append(matches, match)
		}
	}
	return matches
}

func isValidGCODE(code string) bool {
	return (len(code) >= 2) && !strings.Contains(code, ".") && (strings.HasPrefix(code, "G") || strings.HasPrefix(code, "M"))
}

// ShowSummary prints a summary to stdout after scraping.
func ShowSummary(start time.Time, count int) {
	fmt.Println(
		Sprintf("%s %d GCODE commands found in %2.2f seconds and saved to the database.",
			Inverse("[DONE]").Bold(),
			Green(count).Bold(),
			Green(time.Now().Sub(start).Seconds()).Bold()))
}
