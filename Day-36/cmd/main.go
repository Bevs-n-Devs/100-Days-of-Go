package main

import "fmt"

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nGRrequest Handling in Go")

	intro := `
In this lesson, you'll learn how to use GRequests, a Go library inspired by Python's popular requests module. 
GRequests simplifies making HTTP requests by providing a higher-level, more user-friendly API while still leveraging Go's concurrency and efficiency.
	`
	fmt.Println(intro)

	topics := `
	Topics Covered:

1. Introduction to GRequests
- Overview of GRequests as a library inspired by Python's requests module.
- Installation and setup of GRequests in a Go project:

	- Install GRequests using 'go get github.com/levigross/grequests'

2. Basic HTTP Requests
- Performing a GET request with an example.
- Performing a POST request with a JSON payload.

3. Advanced Features
- Adding headers and query parameters to requests.
- Handling timeouts for HTTP requests.

4. File Handling
- Uploading files using FileUpload.

5. Concurrency with GRequests
- Making multiple HTTP requests concurrently using goroutines and sync mechanisms.
	`
	fmt.Println(topics)

	basicUsageGRequests()

	headersAndQueryParameters()

	fileUpload()
}
