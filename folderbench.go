package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	version := "v1.0"

	fmt.Println("You have started up folderbench " + version + ".")
	fmt.Println()

	askFolderCount()
}

func askFolderCount() int {
	foldercount := 0
	askAgain := true
	for askAgain {
		foldercount_str := ask("How many folders? (type exit to exit): ")
		if foldercount_str == "exit" {
			os.Exit(0)
		}
		if _, err := strconv.ParseInt(foldercount_str, 10, 64); err != nil {
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
