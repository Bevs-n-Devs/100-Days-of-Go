/*
set up the database connection function and initialize the database/sql connection pool.
*/

// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nCRUD API with Go and PostgreSQL")

	if err := initDB(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getUsers(w, r)
		case http.MethodPost:
			createUser(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getUserByID(w, r)
		case http.MethodPut:
			updateUser(w, r)
		case http.MethodDelete:
			deleteUser(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*
You’ve now built a CRUD API with a PostgreSQL database, which is fundamental for full-stack development and backend API creation.
This example also covers data persistence, concurrency handling, and database interaction, which are essential skills for Go
developers working on production-ready applications.
*/
