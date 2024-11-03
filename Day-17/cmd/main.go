package main

import "fmt"

func main() {
	fmt.Println("Hello world, hello Yaw!")
	fmt.Println("\nMiddleware in Go for HTTP Servers")

	executeLoggingMiddleware()

	executeAuthMiddleware()
}

/*
Middleware is a crucial concept for building production-ready HTTP servers, as it allows you to intercept and modify requests/responses.
Middleware functions are typically used for tasks like logging, authentication, rate limiting, or adding headers. In Go, middleware is
implemented by wrapping HTTP handlers with additional functionality.

Goals:
- Understand what middleware is and why it’s useful.
- Create a basic logging middleware.
- Implement an authentication middleware for secure routes.

In Go, middleware functions take in an http.Handler (or http.HandlerFunc), add additional processing before and/or after the handler
executes, and then pass the request down the chain.

-- Middleware function signature
func middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        -- Middleware logic before the request reaches the handler
        next.ServeHTTP(w, r)
        -- Middleware logic after the handler has responded
    })
}

*/
