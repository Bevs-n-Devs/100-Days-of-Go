package main

import "fmt"

func Example6() {
	fmt.Println("\nExample 6: Predictable Error Handling")
	var scenario = `
	Violation of PoLA:
	// Function returns nil on error but doesn't document it
	func GetUser(id int) *User {
		if id <= 0 {
			return nil // Surprising to users
		}
		return &User{ID: id, Name: "Example"}
	}

	Following PoLA:
	// Explicit error handling
	func GetUser(id int) (*User, error) {
		if id <= 0 {
			return nil, fmt.Errorf("invalid ID")
		}
		return &User{ID: id, Name: "Example"}, nil
	}
		- Users expect errors to be returned explicitly in Go.
	`
	fmt.Println(scenario)
}
