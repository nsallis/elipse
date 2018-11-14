package log

import (
	"flag"
	"fmt"
)

var logLevel int

// LogLevels allows us to ignore anything over a specific log level
var LogLevels = map[string]int{
	"error": 0,
	"warn":  1,
	"info":  2,
	"debug": 3,
}

// InitLogs initializes log levels based on arguments
func InitLogs() int {
	if logLevel < 1 {
		usage := "level of logging to show"
		defaultLevel := "debug"
		var level string
		flag.StringVar(&level, "log_level", defaultLevel, usage)
		flag.StringVar(&level, "ll", defaultLevel, usage)
		flag.Parse()
		fmt.Println("Using log level: ", level)
		logLevel = LogLevels[level]
		return LogLevels[level]
	}
	return logLevel
}
