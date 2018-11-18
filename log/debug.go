package log

import (
	"fmt"
	"github.com/fatih/color"
	"log"
)

// Debug logs a debug message
func Debug(debugText string, a ...interface{}) {
	if logLevel >= 3 {
		debugText = fmt.Sprintf(debugText, a)
		red := color.New(color.FgHiMagenta).SprintFunc()
		var debugMessage string
		debugMessage = fmt.Sprintf("%s\n\n", debugText)
		log.SetPrefix(red("[DEBUG] "))
		log.Print(debugMessage)
	}
}
