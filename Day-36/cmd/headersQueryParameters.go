package main

import (
	"fmt"
	"log"

	"github.com/levigross/grequests"
)

func headersAndQueryParameters() {
	// request options with headers and query parameters
	rObject := &grequests.RequestOptions{
		Headers: map[string]string{
			"Authorization": "Bearer yout_token",
		},
		Params: map[string]string{
			"search": "golang",
			"page":   "1,",
		},
	}

	// perform a GET request with headers and query parameters
	response, err := grequests.Get("https://api.example.com/resource", rObject)
	if err != nil {
		log.Fatalf("Failed to make GET request: %v", err)
	}

	// print response
	fmt.Println("\nGET request with headers and query parameters")
	fmt.Println("Response body:", response.String())

}
