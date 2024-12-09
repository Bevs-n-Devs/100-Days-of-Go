package main

import "fmt"

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\n	Heimdall For Scalable & Resilient HTTP Calls")

	intro := `
We will dive into Heimdall, a powerful and flexible Go library used for making resilient HTTP calls. 
Heimdall helps developers implement patterns such as retries, timeouts, circuit breakers, and more 
to build robust and fault-tolerant systems.
	`
	fmt.Println(intro)

	objectives := `
	Objectives:

1. What is Heimdall?
2. Why use Heimdall?
3. Features and Benefits of Heimdall
4. Drawbacks of Heimdall
5. Practical Examples
6. Best Practices for Using Heimdall
	`
	fmt.Println(objectives)

	whatIsHeimdall := `
	What is Heimdall?

Heimdall is a Go library that provides utilities for building resilient and fault-tolerant HTTP clients. 
It is inspired by principles from the Hystrix library for circuit breakers and resilience patterns. 
Heimdall is often used in microservices, distributed systems, or anywhere HTTP calls require resilience 
against failures.
	`
	fmt.Println(whatIsHeimdall)

	whyUseHeimdall := `
	Why Use Heimdall?

When making HTTP requests in production systems, various issues can occur:

	- Network Latency: Delays in response.
	- Server Downtime: The target server may become unavailable.
	- Rate Limiting: Too many requests may trigger throttling.
	- Transient Errors: Temporary failures like timeouts or HTTP 500 errors.
	- Heimdall provides a clean and efficient way to handle these challenges by implementing retry policies, backoff strategies, timeouts, and circuit breakers.	
	`
	fmt.Println(whyUseHeimdall)

	features := `
	Features of Heidmall:

1. Retries with Custom Backoff:
	- Retry failed requests using strategies like constant backoff or exponential backoff.
2. Circuit Breakers:
	- Prevent repeated requests to a failing service.
3. Custom Timeouts:
	- Set time limits for HTTP requests to avoid hanging.
4. HTTP Client Wrapping:
	- Wraps Go’s http.Client for enhanced functionality.
5. Pluggable:
	- Easily integrates into existing codebases.
	`
	fmt.Println(features)

	benefits := `
	Benefits of Heimdall:

1. Improved Resilience:
	- Ensures fault tolerance by retrying transient failures.
2. Reduced Downtime:
	- Circuit breakers stop overloading failing services.
3. Efficient Resource Usage:
	- Prevents excessive retries and wasted compute resources.
4. Easy to Use:
	- Offers clean and customizable APIs.	
	`
	fmt.Println(benefits)

	drawbacks := `
	Drawbacks of Heimdall:

1. Increased Complexity:
	- Adds complexity to simple HTTP calls if overused.
2. Limited Support for Advanced Features:
	- Heimdall does not directly manage rate limiting, unlike some other libraries.
3. Dependency Overhead:
	- Adding Heimdall introduces a new dependency to your project.
	`
	fmt.Println(drawbacks)

	// fmt.Println("\n	Practical Examples")

	// basicHTTPClientWithRetries()

	// circuitBreaker()

	bestPractices := `
	Best Practices for Using Heimdall:

1. Use Heimdall Where Necessary
	- Apply Heimdall for external HTTP calls where failures are likely (e.g., third-party APIs or microservices).
2. Circuit Breaker Configuration
	- Configure sensible failure thresholds and timeout durations to avoid false positives.
3. Exponential Backoff Over Constant Backoff
	- Use exponential backoff for transient errors to avoid overwhelming the server.
4. Logging and Monitoring
	- Log retries and circuit breaker states to monitor system health.
5. Combine with Context
	- Pass a context.Context to cancel requests dynamically when necessary.
	`
	fmt.Println(bestPractices)
}
