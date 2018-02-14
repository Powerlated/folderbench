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

	endTime := int(time.Now().Unix())

	finalTime := endTime - startTime

	foldersPerSecond := foldersCreated / finalTime

	fmt.Println()
	fmt.Println("Folders per second: " + strconv.Itoa(foldersPerSecond))

	evaluate(finalTime)

	clean()

}

func clean() {
	if exists("fbdir") {
		fmt.Println("Cleaning folders, this might take a long time...")
		os.RemoveAll("fbdir")
		fmt.Println("Removed fbdir directory")
	}
}

func evaluate(time int) {

}
