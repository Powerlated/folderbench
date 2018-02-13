package main

import (
	"fmt"
	"os"

	input "github.com/tcnksm/go-input"
)

func main() {
	version := "v1.0"

	fmt.Println("You have started up folderbench " + version + ".")
	fmt.Println()
	askAgain := true
	for askAgain {
		//foldercount := ask("How many folders?: ")

	}

}

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
