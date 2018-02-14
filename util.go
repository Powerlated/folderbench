package main

import (
	"os"

	input "github.com/tcnksm/go-input"
)

// Use an external library for asking questions
func ask(message string) string {
	ui := &input.UI{
		Writer: os.Stdout,
		Reader: os.Stdin,
	}

	query := message
	answer, err := ui.Ask(query, &input.Options{
		Default:  "",
		Required: true,
		Loop:     false,
	})
	_ = err
	return answer
}

// Check if a file or folder exists
func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return false
}
