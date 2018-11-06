package log

import (
	"fmt"
	"github.com/fatih/color"
	"log"
)

// Error logs an error to the log file
func Debug(debugText string) {
	// Hardcode error file for now
	// TODO make error file configurable in the future

	red := color.New(color.FgHiMagenta).SprintFunc()
	var debugMessage string
	debugMessage = fmt.Sprintf("%s\n\n", debugText)
	log.SetPrefix(red("[DEBUG] "))
	log.Print(debugMessage)
}
