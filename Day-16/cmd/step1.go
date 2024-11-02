package main

// to simulate an API, create Task resource strcut with fields for ID, name & status
type Task struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}
