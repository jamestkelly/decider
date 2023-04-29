package utils

import (
	"fmt"
	"github.com/fatih/color"
	"time"
)

var (
	base   = "[%s] %s - %s\n"
	cyan   = color.New(color.FgCyan).SprintFunc()
	green  = color.New(color.FgGreen).SprintFunc()
	yellow = color.New(color.FgYellow).SprintFunc()
	red    = color.New(color.FgRed).SprintFunc()
)

/*
Logger
Helper logger function to improve quality and readability of messages printed to the terminal.
Messages are differentiated into three distinct levels following standard practice. These include:

- INFO
- WARN
- ERROR

These messages are printed to the terminal with a corresponding colour and timestamp.
*/
func Logger(message, severity string) {
	msgTime := time.Now().Format("2006-01-02 15:04:05")

	switch severity {
	case "INFO":
		fmt.Printf(base, msgTime, green(severity), cyan(message))
	case "WARN":
		fmt.Printf(base, msgTime, yellow(severity), cyan(message))
	case "ERROR":
		fmt.Printf(base, msgTime, red(severity), cyan(message))
	}
}
