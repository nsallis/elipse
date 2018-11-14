package log

import (
	"fmt"
	"github.com/fatih/color"
	"log"
)

// Error logs an error to the log file
func Error(errorText string, err error) {
	if logLevel >= 0 { // redundant but leaving for consistancy
		red := color.New(color.FgRed).SprintFunc()
		var errorMessage string
		if err != nil {
			errorMessage = fmt.Sprintf("%s\n%s\n\n", errorText, err.Error())
		} else {
			errorMessage = fmt.Sprintf("%s\n\n", errorText)
		}
		log.SetPrefix(red("[ERROR] "))
		log.Print(errorMessage)
	}
}
