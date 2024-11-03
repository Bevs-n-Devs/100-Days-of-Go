package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

/*
In this example, every request to / or /tasks will be logged with the method, path, and response time
*/

// LoggingMiddleware logs details of each incoming HTTP request
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r) // call the actual handler

		log.Printf("Completed %s %s in %v", r.Method, r.URL.Path, time.Since(start))
	})
}

func executeLoggingMiddleware() {
	fmt.Println("\nLogging Middleware Example")
	// Wrap handlers with middleware
	http.Handle("/", LoggingMiddleware(http.HandlerFunc(homeHandler)))
	http.Handle("/tasks", LoggingMiddleware(http.HandlerFunc(tasksHandler)))

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Task API!")
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Tasks endpoint")
}
