package main

import (
	"fmt"
	"os"
	"strconv"
)

// Defines the default foldercount for use when benchmarking
var defaultFoldercount = 250000
// Version of the program
var version = "vl.0"

func main() {
	fmt.Println("You have started up folderbench " + version + ".")
	fmt.Println()

	// Gets the amount of folders from the user and calls bench() with it
	bench(askFolderCount())	
}

// Asks the user how many folders to create
func askFolderCount() int {
	foldercount := 0
	askAgain := true

	// Loops until a valid answer has been provided
	for askAgain {
		// Use an external library to ask the user with a message
		foldercountStr := ask("How many folders? (type exit to exit): ")
		if foldercountStr == "exit" {
			// Exit the program if "exit" is provided
			fmt.Println("Exiting...")
			os.Exit(0)
		} else if foldercountStr == "" {
			// Return the default foldercount if the input is empty
			fmt.Println("It seems like no number has been provided. Using the default of " + strconv.Itoa(defaultFoldercount) + ".")
			return defaultFoldercount
		} else if _, err := strconv.ParseInt(foldercountStr, 10, 64); err != nil {
			// Rejects non-numbers
			fmt.Println("It seems like that is not a number.")
			fmt.Println()
		} else {
			// A number has been provided, return it.
			askAgain = false
			fmt.Println()
			foldercount, err = strconv.Atoi(foldercountStr)
			_ = err

		}
	}
	return foldercount
}
