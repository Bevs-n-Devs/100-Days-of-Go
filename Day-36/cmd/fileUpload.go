package main

import (
	"fmt"
	"log"
	"os"

	"github.com/levigross/grequests"
)

func fileUpload() {
	// open the file
	file, err := os.Open("example.txt")

	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// create new request options
	ro := &grequests.RequestOptions{
		File: grequests.FileUpload{
			FileName:     "example.txt",
			FileContents: file,
		},
	}

	// perform post request tp upload file
	response, err := grequests.Post("https://example.com/upload", ro)
	if err != nil {
		log.Fatalf("Failed to upload file: %v", err)
	}

	// print response
	fmt.Println("Response status:", response.StatusCode)
	fmt.Println("Response body:", response.String())
}
