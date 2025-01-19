package main

import "fmt"

func Example2() {
	fmt.Println("\nExample 2: Over-Abstraction in Code")

	var scenario = `
	Scenario:
	You need to fetch data from a database.

	Over-Complicated Approach:
	Adding multiple layers of abstraction

	Code Example:
	type Database interface {
		Query(query string) ([]map[string]interface{}, error)
	}

	type MySQLDatabase struct{}

	func (db MySQLDatabase) Query(query string) ([]map[string]interface{}, error) {
		// Simulate database query
		return []map[string]interface{}{
			{"id": 1, "name": "Alice"},
			{"id": 2, "name": "Bob"},
		}, nil
	}

	func FetchData(db Database) ([]map[string]interface{}, error) {
		return db.Query("SELECT * FROM users")
	}

	Problems:
		1. The Database interface and MySQLDatabase implementation are unnecessary for a simple query.
		2. Adds layers without any real advantage.

	KISS Approach:
	Keep it straightforward unless you need flexibility.

	Code Example:
	func FetchData() ([]map[string]interface{}, error) {
		// Simulate database query
		return []map[string]interface{}{
			{"id": 1, "name": "Alice"},
			{"id": 2, "name": "Bob"},
		}, nil
	}

	Benefits:
		- Fewer layers to maintain.
		- Easier to understand and debug.
	`
	fmt.Println(scenario)
}
