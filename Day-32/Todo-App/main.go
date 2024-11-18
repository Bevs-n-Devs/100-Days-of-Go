package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

// create a new todo example data
var todos = []Todo{
	{ID: 1, Task: "Learn Go"},
	{ID: 2, Task: "Learn Docker"},
	{ID: 3, Task: "Learn Kubernetes"},
	{ID: 4, Task: "Learn Python"},
	{ID: 5, Task: "Become a Go Developer"},
	{ID: 6, Task: "Get a new car - electric Jaguar"},
	{ID: 7, Task: "Get a flat in Manchester"},
}

func main() {
	r := gin.Default()

	// serve static files and templates
	r.LoadHTMLGlob("templates/*")

	// define routes related to TODO
	todoRoutes := r.Group("/todos")
	{
		todoRoutes.GET("/", getTodos)    // view all TODOs
		todoRoutes.POST("/", createTodo) // create a new TODO

	}

	// root route to serve the HTML
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", todos)
	})

	log.Println("Server running on http://localhost:8080")
	r.Run(":8080")
}

// get all TODOs
func getTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

// create a new TODO
func createTodo(c *gin.Context) {
	task := c.PostForm("Task") // get the "Task" field from the HTML form
	if task == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task cannot be empty"})
		return // stop the function
	}

	// create a new TODO from the HTML form
	newTodo := Todo{
		ID:   len(todos) + 1,
		Task: task,
	}

	todos = append(todos, newTodo)
	c.Redirect(http.StatusFound, "/")
}
