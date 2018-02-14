package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func bench(foldercount int) {

	clean()

	startTime := int(time.Now().Unix())

	os.Mkdir("fbdir", 0777)
	os.Chdir("fbdir")

	fmt.Println("Benchmarking with " + strconv.Itoa(foldercount) + " folders")

	foldersCreated := 0
	for foldersCreated < foldercount+1 {

		if foldersCreated%10000 == 0 && foldersCreated != 0 {
			fmt.Println("Made " + strconv.Itoa(foldersCreated) + " folders")
		}

		os.Mkdir(strconv.Itoa(foldersCreated), 0777)
		foldersCreated++

	}

	os.Chdir("..")

	endTime := int(time.Now().Unix())

	finalTime := endTime - startTime

	if finalTime == 0 {
		finalTime = 1
	}

	foldersPerSecond := foldersCreated / finalTime

	fmt.Println()
	fmt.Println("Folders per second: " + strconv.Itoa(foldersPerSecond))

	fmt.Println()
	fmt.Println(evaluate(foldersPerSecond))

	clean()

}

func clean() {
	fmt.Println("Cleaning folders, this might take a long time...")

	os.RemoveAll("fbdir")
	fmt.Println("Removed fbdir directory")
}

func evaluate(fps int) string {
	if fps >= 3500 {
		return "You're on Linux, aren't you?"
	} else if fps >= 3000 {
		return "Amazing."
	} else if fps >= 2500 {
		return "Very good."
	} else if fps >= 2000 {
		return "It's a great filesystem if it can handle this many folders being created."
	} else if fps >= 1500 {
		return "It's a decently good filesystem."
	} else if fps >= 1000 {
		return "Absolute garbage, just absolute garbage. How can your storage or filesystem be so slow?"
	}
	return "Contact the developer, you shouldn't see this line of text."
}
