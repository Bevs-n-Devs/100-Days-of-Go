package main

import "fmt"

func Factor7() {
	var factor7 = `
	7. Port Binding: Export Services via Port Binding
		- Principle: The app should expose services (eg. HTTP) by binding to a port
		- Why: Keeps apps self-contained and decoupled from external routing.

		Code Example:
			func main() {
				r := gin.Default()
				r.GET("/", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "Hello, World!"})
				})
				r.Run(":8080") // Binds to port 8080
			}
	`
	fmt.Println(factor7)
}
