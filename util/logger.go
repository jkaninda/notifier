package util

import (
	"fmt"
	"os"
	"time"
)

var currentTime = time.Now().Format("2006/01/02 15:04:05")

// Info message
func Info(msg string, args ...any) {
	formattedMessage := fmt.Sprintf(msg, args...)
	if len(args) == 0 {
		fmt.Printf(msg)
	} else {
		fmt.Printf(formattedMessage)
	}
}

// Warn Warning message
func Warn(msg string, args ...any) {
	formattedMessage := fmt.Sprintf(msg, args...)
	if len(args) == 0 {
		fmt.Printf(msg)
	} else {
		fmt.Printf(formattedMessage)
	}
}

// Error error message
func Error(msg string, args ...any) {
	formattedMessage := fmt.Sprintf(msg, args...)
	if len(args) == 0 {
		fmt.Printf(msg)
	} else {
		fmt.Printf(formattedMessage)
	}
}
func Done(msg string, args ...any) {
	formattedMessage := fmt.Sprintf(msg, args...)
	if len(args) == 0 {
		fmt.Printf(msg)
	} else {
		fmt.Printf(formattedMessage)
	}
}

// Fatal logs an error message and exits the program.
func Fatal(msg string, args ...any) {
	formattedMessage := fmt.Sprintf(msg, args...)
	if len(args) == 0 {
		fmt.Printf(msg)
	} else {
		fmt.Printf(formattedMessage)
	}
	os.Exit(1)
}
