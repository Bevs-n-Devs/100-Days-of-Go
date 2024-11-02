package main

import (
	"encoding/json"
	"net/http"
)

/*
functions to implement CRUD operations
*/

// get all tasks
func getAllTasks(w http.ResponseWriter, r *http.Request) {
	tasksMutex.Lock()
	defer tasksMutex.Unlock()
	json.NewEncoder(w).Encode(tasks)
}

// create a new task
func createTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	tasksMutex.Lock()
	task.ID = nextID
	nextID++
	tasks = append(tasks, task)
	tasksMutex.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

// get task by ID
func getTaskByID(w http.ResponseWriter, id int) {
	tasksMutex.Lock()
	defer tasksMutex.Unlock()
	for _, task := range tasks {
		if task.ID == id {
			json.NewEncoder(w).Encode(task)
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}

// update a task
func updateTask(w http.ResponseWriter, r *http.Request, id int) {
	var updateTask Task
	if err := json.NewDecoder(r.Body).Decode(&updateTask); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	tasksMutex.Lock()
	defer tasksMutex.Unlock()
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Name = updateTask.Name
			tasks[i].Completed = updateTask.Completed
			json.NewEncoder(w).Encode(tasks[i])
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}

// delete a task
func deleteTask(w http.ResponseWriter, id int) {
	tasksMutex.Lock()
	defer tasksMutex.Unlock()
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}
