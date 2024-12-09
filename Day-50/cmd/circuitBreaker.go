package main

import (
	"fmt"
	"time"

	"github.com/gojek/heimdall/v7/circuitbreaker"
	"github.com/gojek/heimdall/v7/httpclient"
)

func circuitBreaker() {
	fmt.Println("\nCircuit Breaker")

	// Create a circuit breaker with a failure threshold of 3
	cb := circuitbreaker.NewCircuitBreaker("my-breaker")
	client := httpclient.NewClient(
		httpclient.WithHTTPTimeout(5*time.Second),
		httpclient.WithRetryCount(3),
		httpclient.WithRetrier(httpclient.NewRetrier(
			httpclient.NewExponentialBackoff(500*time.Millisecond, 5*time.Second))),
		httpclient.WithCircuitBreaker(cb),
	)

	// Perform an HTTP GET request
	res, err := client.Get("https://jsonplaceholder.typicode.com/posts/1", nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Response:", res.StatusCode)
}
