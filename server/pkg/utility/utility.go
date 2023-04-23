package utility

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

func logger(severity string, message string) {
	// Set color based on severity level
	var colorFunc func(...interface{}) string
	switch severity {
	case "INFO":
		colorFunc = color.New(color.FgGreen).SprintFunc()
	case "WARN":
		colorFunc = color.New(color.FgYellow).SprintFunc()
	case "ERROR":
		colorFunc = color.New(color.FgRed).SprintFunc()
	default:
		colorFunc = color.New(color.FgWhite).SprintFunc()
	}

	// Format log message with current time and severity level
	logMessage := fmt.Sprintf("[%s] %s - %s", time.Now().Format("2006-01-02 15:04:05"), severity, message)

	// Print to console with color
	fmt.Fprintln(os.Stdout, colorFunc(logMessage))
}
