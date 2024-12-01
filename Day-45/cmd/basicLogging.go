package main

import (
	"fmt"
	"log"
	"net/http"
)

// middleware function
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r) // call the next handler
	})
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the main page!"))
}

func basicLoggingMiddleware() {
	fmt.Println("\nBasic Logging Middleware")

	max := http.NewServeMux()

	// wrap the main handler with the middleware
	max.Handle("/", LoggingMiddleware(http.HandlerFunc(mainHandler)))

	log.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", max)
}

/*
What’s Happening?
	1. The LoggingMiddleware function wraps any http.Handler.
	2. It logs the request method and path.
	3. The next.ServeHTTP call passes the request to the next handler in the chain.
*/
