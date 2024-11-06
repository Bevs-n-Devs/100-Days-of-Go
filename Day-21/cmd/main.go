package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
Build a REST API in Go that includes authentication via JSON Web Tokens (JWT).
Adding authentication will help protect routes and ensure only authenticated users can access certain resources.

- JWT (JSON Web Tokens): A secure way to transmit information between parties as a JSON object.
						 This token will contain claims, such as a user’s ID, and will be signed to verify authenticity.

- Middleware: Middleware functions intercept requests to check for authentication before granting access to protected routes.
*/

func main() {
	fmt.Println("Hello world, hello Yaw!")
	fmt.Println("\nREST API with JWT Authentication")

	// set up the HTTP server and use the middleware to protect the /profile route
	http.HandleFunc("/register", RegisterHandler)
	http.HandleFunc("/login", LoginHandler)

	// protected route
	http.Handle("/profile", AuthMiddleware(http.HandlerFunc(ProfileHandler)))

	log.Println("\nServer running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*
TESTING THE API

Register User
curl -X POST -d '{"username": "user1", "password": "password123"}' -H "Content-Type: application/json" http://localhost:8080/register

Login to get JWT
curl -X POST -d '{"username": "user1", "password": "password123"}' -H "Content-Type: application/json" http://localhost:8080/login

Access Protected Route
curl -X GET -H "Authorization: Bearer your_jwt_token" http://localhost:8080/profile

*/
