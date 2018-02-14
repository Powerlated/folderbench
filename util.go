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

func exists(path string) (bool) {
	if _, err := os.Stat(path); err != nil {
		if os.IsExist(err) {
			return true;
		} else {
			return false;
		}
	}
	return false;
}