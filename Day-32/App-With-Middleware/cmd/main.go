package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware(c *gin.Context) {
	// log the incoming request method and path
	method := c.Request.Method
	path := c.Request.URL.Path
	c.Next() // continue to the next handler or middleware
	// log th eresponse status code after proceeding
	status := c.Writer.Status()
	log.Printf("Method: %s, Path: %s, Status: %d", method, path, status)
}

func main() {
	r := gin.Default()

	// apply the LoggerMiddleware gloabbly (to all requests)
	r.Use(LoggerMiddleware)

	// define some routes
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Middleware with Gin!"})
	})

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello world!"})
	})

	r.GET("/goodbye", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Goodbye world!"})
	})

	log.Println("Server running on http://localhost:8080")
	r.Run(":8080")
}
