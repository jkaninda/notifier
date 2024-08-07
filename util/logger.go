package util

import (
	"fmt"
	"os"
)

// Info message
func Info(msg string, args ...any) {
	formattedMessage := fmt.Sprintf(msg, args...)
	if len(args) == 0 {
		fmt.Printf("%s\n", msg)
	} else {
		fmt.Printf("%s\n", formattedMessage)
	}
}

// Fatal logs an error message and exits the program.
func Fatal(msg string, args ...any) {
	formattedMessage := fmt.Sprintf(msg, args...)
	if len(args) == 0 {
		fmt.Printf("%s\n", msg)
	} else {
		fmt.Printf("%s\n", formattedMessage)
	}
	os.Exit(1)
}
