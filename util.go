package main

import (
	"os"

	input "github.com/tcnksm/go-input"
)

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
