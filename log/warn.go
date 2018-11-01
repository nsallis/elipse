package log

import (
	"fmt"
	"github.com/fatih/color"
	"log"
)

// Warn logs an error to the log file
func Warn(warnText string, err error) {
	// Hardcode error file for now
	// TODO make error file configurable in the future
	yellow := color.New(color.FgYellow).SprintFunc()
	var warnMessage string
	if err != nil {
		warnMessage = fmt.Sprintf("%s\n%s\n\n", warnText, err.Error())
	} else {
		warnMessage = fmt.Sprintf("%s\n\n", warnText)
	}
	log.SetPrefix(yellow("[WARN] "))
	log.Print(warnMessage)
}
