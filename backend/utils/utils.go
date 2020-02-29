package utils

import (
	"fmt"
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

func ShowSummary(start time.Time, count int) {
	fmt.Println(
		Sprintf("%s %d GCODE commands found in %2.2f minutes and saved to the database.",
			Inverse("[DONE]").Bold(),
			Green(count).Bold(),
			Green(time.Now().Sub(start).Minutes()).Bold()))
}
