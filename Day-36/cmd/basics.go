package main

import (
	"fmt"
	"log"

	"github.com/levigross/grequests"
)

// GRequests provides simple functions to perform GET, POST, PUT, DELETE, etc. Here are some examples.
func basicUsageGRequests() {
	fmt.Println("\nBasic Usage of GRequests")

	// perform a GET request
	res, err := grequests.Get("https://jsonplaceholder.typicode.com/posts/1", nil)
	if err != nil {
		log.Fatalf("Failed to make GET request: %v", err)
	}

	// print response
	fmt.Println("\nGET request")
	fmt.Println("Response status:", res.StatusCode)
	fmt.Println("Response body:", res.String())

	// JSON payload
	data := map[string]string{
		"title":  "foo",
		"body":   "bar",
		"userId": "1",
	}

	// create a request options object with JSON data
	rObject := &grequests.RequestOptions{
		JSON: data,
	}

	// perform a POST request
	response, err := grequests.Post("https://jsonplaceholder.typicode.com/posts", rObject)
	if err != nil {
		log.Fatalf("Failed to make POST request: %v", err)
	}

	// print the response
	fmt.Println("\nPOST request")
	fmt.Println("Response status:", response.StatusCode)
	fmt.Println("Response body:", response.String())
}
