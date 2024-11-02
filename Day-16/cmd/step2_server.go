package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
The net/http package provides the basics for setting up a server.
Start by creating a simple handler function and configuring the server.
*/

func server() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the Task API!")
		fmt.Fprintln(w, "Endpoints Available:")
		fmt.Fprintln(w, "/tasks - GET to retrieve all tasks, POST to create a new task.")
		fmt.Fprintln(w, "/tasks/{id} - GET, PUT or DELETE a task by ID")
	})

	http.HandleFunc("/tasks", tasksHandler)
	http.HandleFunc("/tasks/", taskHandler)

	log.Print("Sever running on localhost - http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
