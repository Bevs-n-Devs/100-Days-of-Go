package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// create new Gin router
	r := gin.Default()

	// define GET route
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Gin!",
		})
	})

	// define another GET route
	r.GET("/hello/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello world, hello " + name,
		})
	})

	// start the server on port 8080
	r.Run(":8080") // default is 8080
}
