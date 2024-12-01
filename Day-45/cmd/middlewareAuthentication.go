package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// middleware function
func AuthMiddleare() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "my-secret-token" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort() // abort the request
			return
		}
	}
}

func customMiddlewareForAuthentication() {
	fmt.Println("\nCustom Middleware for Authentication")

	r := gin.Default()

	// apply middleware globally
	r.Use(AuthMiddleare())

	r.GET("/secure", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "You have access!"})
	})

	r.GET("/public", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "This is public"})
	})

	r.Run(":8080")
}

/*
What’s Happening?
	1. The AuthMiddleware function checks for a specific token in the request header.
	2. If the token is missing or incorrect, it aborts the request.
	3. If valid, it calls c.Next() to allow the request to proceed.
*/
