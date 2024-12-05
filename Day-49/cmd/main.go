package main

import (
	"fmt"

	"github.com/Bevs-n-Devs/100-Days-of-Go/Day-49/cmd/handlers"
	"github.com/Bevs-n-Devs/100-Days-of-Go/Day-49/cmd/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\n	Building RESTful APIs With Authentication in Go")

	intro := `
In this lesson, we’ll take a deep dive into building RESTful APIs with a focus on implementing authentication. 
Authentication is a critical feature for securing APIs, especially in scenarios where sensitive data or actions 
are involved. We will explore token-based authentication using JSON Web Tokens (JWT) and best practices for securing APIs.
	`
	fmt.Println(intro)

	objectives := `
	Objectives

1. Understand what RESTful APIs are.
2. Learn about the importance of authentication in APIs.
3. Implement JWT-based authentication in a RESTful API using Go.
4. Follow best practices for securing and structuring APIs.
	`
	fmt.Println(objectives)

	whatIsRestAPI := `
	What Are RESTful APIs?

REST (Representational State Transfer) APIs follow a set of principles for designing networked applications. RESTful APIs use HTTP methods to interact with resources:

GET: Retrieve a resource.
POST: Create a new resource.
PUT: Update an existing resource.
DELETE: Remove a resource.
APIs are stateless, meaning each request must contain all necessary information.
	`
	fmt.Println(whatIsRestAPI)

	essentialAuthentication := `
	Why Authentication is Essential

Authentication ensures that only authorized users can access specific resources or perform actions. 
Token-based authentication is one of the most popular methods, using tokens like JWT to validate users.
	`
	fmt.Println(essentialAuthentication)

	toolsAndLibraries := `
	Tools and Libraries

To build and secure our RESTful API, we’ll use:

	- Gin: A web framework for building APIs.
	- JWT-Go: A library for working with JSON Web Tokens.

Install the required libraries:

	go get -u github.com/gin-gonic/gin
	go get -u github.com/golang-jwt/jwt/v5
	`
	fmt.Println(toolsAndLibraries)

	r := gin.Default()

	// public route
	r.POST("/login", handlers.Login)

	// protected route - implement middleware
	auth := r.Group("/user")
	auth.Use(middleware.AuthNMiddleware())
	auth.GET("/profile", handlers.GetProfile)

	r.Run(":8080")

}
