package log

import (
	"fmt"
	"github.com/fatih/color"
	"log"
)

// Info logs an error to the log file
func Info(infoText string) {
	// Hardcode error file for now
	// TODO make error file configurable in the future
	green := color.New(color.FgGreen).SprintFunc()
	infoMessage := fmt.Sprintf("%s\n\n", infoText)
	log.SetPrefix(green("[INFO] "))
	log.Print(infoMessage)
}
