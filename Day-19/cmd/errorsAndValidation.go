package main

import (
	"encoding/json"
	"fmt"
	"log"
)

/*
Handling Invalid JSON Fields Gracefully
When dealing with user input, fields might be missing or malformed. Here’s how you can handle such cases:
*/

func decodeAndValidateJSON(data []byte) (Task, error) {
	var task Task
	err := json.Unmarshal(data, &task)
	if err != nil {
		return task, fmt.Errorf("error decoding JSON: %v", err)
	}

	if task.Title == "" {
		return task, fmt.Errorf("validation error: title cannot be empty")
	}

	return task, nil
}

func errorsValidation() {
	fmt.Println("\nErrors and Validation")
	jsonData := `{"id": 2, "completed": false}`

	task, err := decodeAndValidateJSON([]byte(jsonData))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("Valid Task: %+v\n", task)
}
