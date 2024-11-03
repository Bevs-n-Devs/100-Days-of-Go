package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
A simple token-based authentication middleware. For demonstration,
we’ll use a hardcoded token (in real applications, you'd use a more secure approach).
*/

// AuthMiddleware checks for a valid token in the "Authorization" header
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "Bearer mysecrettoken" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		// Pass to the next handler if authorized
		next.ServeHTTP(w, r)
	})
}

func executeAuthMiddleware() {
	fmt.Println("\nAuthentication Middleware Example")
	// Wrap `/tasks` with both logging and authentication middlewares
	http.Handle("/", LoggingMiddleware(http.HandlerFunc(homeHandler)))
	http.Handle("/tasks", LoggingMiddleware(AuthMiddleware(http.HandlerFunc(tasksHandler))))

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
