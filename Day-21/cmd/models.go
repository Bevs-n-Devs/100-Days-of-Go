package main

// define user struct and initializing JWT
type User struct {
	ID       int    `json: "id"`
	Username string `json: "username"`
	Password string `json: "password"`
}
