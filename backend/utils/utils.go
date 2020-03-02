package utils

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	. "github.com/logrusorgru/aurora"
)

const debug = true

// Debug TODO...
func Debug(msg string) {
	if debug {
		fmt.Println(Sprintf("%s   %s", White("[DEBUG]").BgGray(8-1), Blue(msg)))
	}
}

// CheckError TODO..
func CheckError(err error) {
	if err != nil {
		fmt.Println(Sprintf("%s   %s", White("[ERROR]").BgRed(), White(err)))
	}
}

// Find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
// https://golangcode.com/check-if-element-exists-in-slice/
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// FindGCODEInString returns all instances of valid commands within any string.
func FindGCODEInString(str string) []string {
	// Filter out CNC commands.
	if strings.Contains(str, "CNC") {
		return nil
	}

	// Find the GCODE command(s) in the string via regex.
	var matches []string
	var re = regexp.MustCompile(`(?m)[MG]\d{1,3}`)
	for _, match := range re.FindAllString(str, -1) {
		matches = append(matches, match)
	}
	return matches
}

// ShowSummary prints a summary to stdout after scraping.
func ShowSummary(start time.Time, count int) {
	fmt.Println(
		Sprintf("%s %d 3d printer GCODE commands found in %2.2f seconds. Records saved in database.",
			Inverse("[DONE]").Bold(),
			Green(count).Bold(),
			Green(time.Now().Sub(start).Seconds()).Bold()))
}
