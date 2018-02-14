package main

import (
	"fmt"
	"os"
	"strconv"
)

var defaultFoldercount = 250000
var version = "vl.0"

func main() {
	fmt.Println("You have started up folderbench " + version + ".")
	fmt.Println()

	bench(askFolderCount())	
}
func askFolderCount() int {
	foldercount := 0
	askAgain := true
	for askAgain {
		foldercountStr := ask("How many folders? (type exit to exit): ")
		if foldercountStr == "exit" {
			fmt.Println("Exiting...")
			os.Exit(0)
		} else if foldercountStr == "" {
			fmt.Println("It seems like no number has been provided. Using the default of " + strconv.Itoa(defaultFoldercount) + ".")
			return defaultFoldercount
		} else if _, err := strconv.ParseInt(foldercountStr, 10, 64); err != nil {
			fmt.Println("It seems like that is not a number.")
			fmt.Println()
		} else {
			askAgain = false
			fmt.Println()
			foldercount, err = strconv.Atoi(foldercountStr)
			_ = err

		}
	}
	return foldercount
}
