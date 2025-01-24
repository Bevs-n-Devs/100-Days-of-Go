package main

import "fmt"

func Example2() {
	fmt.Println("\nExample 2: API Design")
	var scenario = `
	Violation of PoLA:
		- An HTTP endpoint behaves differently from common conventions.
	
	Code Example:
	GET /delete-user?id=123
	
	Problems:
		- Users expect a GET request to be safe and idempotent (not modifying data). Using it to delete a user is unexpected.

	Following PoLA:
	DELETE /users/123
		- The use of the HTTP DELETE method aligns with RESTful conventions and user expectations.
	`
	fmt.Println(scenario)
}
