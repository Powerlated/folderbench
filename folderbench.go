package main

import (
	"fmt"
	"os"
	"strconv"
)

var defaultFoldercount = 10000

func main() {
	version := "vl.0"

	fmt.Println("You have started up folderbench " + version + ".")
	fmt.Println()

	foldercount := askFolderCount()
	bench(foldercount)
}

func askFolderCount() int {
	foldercount := 0
	askAgain := true
	for askAgain {
		foldercount_str := ask("How many folders? (type exit to exit): ")
		if foldercount_str == "exit" {
			fmt.Println("Exiting...")
			os.Exit(0)
		} else if foldercount_str == "" {
			fmt.Println("It seems like no number has been provided. Using the default of " + strconv.Itoa(defaultFoldercount) + ".")
			return defaultFoldercount
		} else if _, err := strconv.ParseInt(foldercount_str, 10, 64); err != nil {
			fmt.Println("It seems like that is not a number.")
			fmt.Println()
		} else {
			askAgain = false
			fmt.Println()
			foldercount, err = strconv.Atoi(foldercount_str)
			_ = err

		}
	}
	return foldercount
}
