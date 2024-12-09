package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/gojek/heimdall/v7/httpclient"
	"github.com/gojek/heimdall/v7/plugins"
)

func basicHTTPClientWithRetries() {
	fmt.Println("\nBasic HTTP Client With Retries")

	// create a constant backoff plugin with 2 seconds between retries and max timeout of 5 seconds
	backoff := plugins.NewConstantBackoff(2*time.Second, 5*time.Second)

	// create a Heimdall HTTP client with retries
	client := httpclient.NewClient(
		httpclient.WithHTTPTimeout(10*time.Second), // Request timeout
		httpclient.WithRetryCount(3),               // Retry up to 3 times
		httpclient.WithRetrier(backoff),            // Use constant backoff
	)

	// perform the HTTP request
	res, err := client.Get("https://jsonplaceholder.typicode.com/posts/1", nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response:", string(body))
}
