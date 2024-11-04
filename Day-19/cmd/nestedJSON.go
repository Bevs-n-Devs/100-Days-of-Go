package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// define structs for JSON data
type Owner struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Task struct {
	ID        int      `json:"id"`
	Title     string   `json:"title"`
	Completed bool     `json:"completed"`
	Owner     Owner    `json:"owner"`
	Tags      []string `json:"tags"`
}

// decode JSON to structs
func decodeJSON() {
	fmt.Println("\nNested JSON: Decoding JSON into structs")

	jsonData := `{
		"id": 1,
		"title": "Learn Go",
		"completed": false,
		"owner": {
			"name": "Alice",
			"email": "alice@email.com",
		},
		"tags": ["programmer", "go", "golang"],
	}`

	var task Task
	err := json.Unmarshal([]byte(jsonData), &task)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	fmt.Printf("Decoded Task: %v\n", task)
}

/*
This will decode the JSON into the Task struct, handling nested structures (like Owner) and arrays (like Tags) automatically.
*/
