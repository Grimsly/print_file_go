package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please enter a filepath")
		os.Exit(1)
	}

	file := os.Args[1]

	// The File struct contains a function that implements the Reader interface
	fileReader, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	io.Copy(os.Stdout, fileReader)
}
