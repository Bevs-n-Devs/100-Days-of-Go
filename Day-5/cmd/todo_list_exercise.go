package main

import "fmt"

type Task struct {
	name string
	done bool
}

func addTask(tasks *[]Task, name string) {
	// create a new task and append it to task list
	*tasks = append(*tasks, Task{name: name, done: false})
}

func markDone(tasks *[]Task, index int) {
	if index >= 0 && index < len(*tasks) {
		(*tasks)[index].done = true
	}
}

func listTasks(tasks []Task) {
	for i, task := range tasks {
		status := "Incomplete"
		if task.done {
			status = "Done"
		}
		fmt.Printf("%d: %s - %s\n", i, task.name, status)
	}
}

func todoListExercise() {
	tasks := []Task{}

	addTask(&tasks, "Complete Go pointers lesson")
	addTask(&tasks, "Build a to-do-app")

	fmt.Println("Tasks before markign done:")
	listTasks(tasks)

	markDone(&tasks, 0)

	fmt.Println("Tasks after markign done:")
	listTasks(tasks)
}
