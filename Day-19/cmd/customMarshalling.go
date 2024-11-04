package main

import (
	"encoding/json"
	"fmt"
	"log"
)

/*
Sometimes you need to modify the JSON format, handle missing fields, or validate data before marshaling/unmarshaling.

Custom Marshaling
Suppose you want to convert Task.Completed from bool to "yes" or "no" in the JSON representation.
*/
func (t Task) MarshallJSON() ([]byte, error) {
	type Alias Task // create an to avoid recursion
	alias := Alias(t)

	// custom modification
	if t.Completed {
		alias.Completed = true
	} else {
		alias.Completed = false
	}

	return json.Marshal(alias)
}

func customMarshalling() {
	fmt.Println("\nCustom Marshalling")

	task := Task{
		ID:        1,
		Title:     "Learn Go",
		Completed: true,
		Owner:     Owner{Name: "Alice", Email: "alice@example.com"},
		Tags:      []string{"programming", "go"},
	}

	jsonData, err := json.Marshal(task)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}

	fmt.Printf("Custom Marshaled JSON: %s\n", jsonData)
}
