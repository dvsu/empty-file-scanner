package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Please provide a directory")
		return
	}

	files, err := ioutil.ReadDir(args[0])
	// if directory does not exist, err != nil and the code block below will be executed
	if err != nil {
		fmt.Println(err)
		return
	}

	// Estimate capacity by calculating the total length
	// of file names in a directory and preallocate the capacity to new slice.
	// This memory optimization is done to prevent the re-creation of new backing
	// array every time the capacity is full.
	var size int

	for _, file := range files {
		// if it is a directory, skip it and go to the next one
		if file.IsDir() {
			continue
		}

		if file.Size() == 0 {
			size += len(file.Name()) + 1 // including the newline character
		}
	}

	// create a slice that will store the names of empty file
	fmt.Printf("Allocate %d bytes in memory...\n", size)
	names := make([]byte, 0, size)

	var count int

	for _, file := range files {
		// if it is a directory, skip it and go to the next one
		if file.IsDir() {
			continue
		}

		if file.Size() == 0 {
			count += 1
			names = append(names, file.Name()...)
			names = append(names, '\n')
		}
	}

	// write the names of empty file to a file named "summary.txt"
	err = ioutil.WriteFile("summary.txt", names, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Printf("%s\n", names)
	content, err := ioutil.ReadFile("summary.txt")

	noun := "files"

	if count == 1 {
		noun = "file"
	}

	fmt.Printf(`
Scan completed!
%-2d empty %s found in %-10s
Result stored in 'summary.txt'    
File size: %4d bytes      
`, count, noun, args[0], len(content))
}
