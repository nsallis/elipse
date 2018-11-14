package log

import (
	"fmt"
	"github.com/fatih/color"
	"log"
)

// Warn logs a warning to the log file
func Warn(warnText string, err error) {
	if logLevel >= 1 {
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
}
