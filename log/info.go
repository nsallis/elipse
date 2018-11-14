package log

import (
	"fmt"
	"github.com/fatih/color"
	"log"
)

// Info logs an info message
func Info(infoText string) {
	if logLevel >= 2 {
		green := color.New(color.FgGreen).SprintFunc()
		infoMessage := fmt.Sprintf("%s\n\n", infoText)
		log.SetPrefix(green("[INFO] "))
		log.Print(infoMessage)
	}
}
