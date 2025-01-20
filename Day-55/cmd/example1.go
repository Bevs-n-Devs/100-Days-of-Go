package main

import "fmt"

func Example1() {
	fmt.Println("\nExample 1: Basic Separation in a Web Application")

	var scenario = `
	Imagine you're building a web application with:
		1. Database access
		2. Buisness logic
		3. HTTP request handling

	Without Separation of Concerns:
		- Here all the logic is in the HTTP handler
	
	Code Example:
	func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
		// Parse request
		username := r.URL.Query().Get("username")
		email := r.URL.Query().Get("email")

		// Business logic
		if username == "" || email == "" {
			http.Error(w, "Missing parameters", http.StatusBadRequest)
			return
		}

		// Database access
		db, _ := sql.Open("mysql", "user:password@/dbname")
		defer db.Close()

		_, err := db.Exec("INSERT INTO users (username, email) VALUES (?, ?)", username, email)
		if err != nil {
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return
		}

		// Respond
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("User created"))
	}

	Problems:
		1. The HTTP handler does too much (request parsing, buisness logic, and database access).
		2. Hard to test individual concerns (eg. buisness logic or database separately).
		3. Changing one concern (eg. the database) risks breaking others.

	With Separation of Concerns
	Split the concerns into layers:
		- Handler: Deals with the HTTP requests and responses.
		- Service: Handles the buisness logic.
		- Repository: Manages database access.

	Code Example:
	// Repository: Handles database operations
	type UserRepository struct {
		DB *sql.DB
	}

	func (repo *UserRepository) CreateUser(username, email string) error {
		_, err := repo.DB.Exec("INSERT INTO users (username, email) VALUES (?, ?)", username, email)
		return err
	}

	// Service: Contains business logic
	type UserService struct {
		Repo *UserRepository
	}

	func (service *UserService) CreateUser(username, email string) error {
		if username == "" || email == "" {
			return fmt.Errorf("username and email are required")
		}
		return service.Repo.CreateUser(username, email)
	}

	// Handler: Manages HTTP requests
	func CreateUserHandler(service *UserService) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			username := r.URL.Query().Get("username")
			email := r.URL.Query().Get("email")

			err := service.CreateUser(username, email)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("User created"))
		}
	}

	Benefits:
		1. Modularity: Each part (handelr, service, repository) has a single responsibility
		2. Reusability: The UserService can be reused in CLI tools or batch jobs.
		3. Testability: You can test each layer independently (eg. mock the respository for testing the service).
	`
	fmt.Println(scenario)

}
