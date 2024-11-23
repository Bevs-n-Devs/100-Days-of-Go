package main

import (
	"fmt"
	"sync"
	"time"
)

func waitGroupsWithAnonymousFunctions() {
	// Using WitGroups with anonymous functions
	fmt.Println("\nUsing WaitGroups with Anonymous Functions")

	var wg sync.WaitGroup

	tasks := []string{"Task 1", "Task 2", "Task 3", "Task 4", "Task 5"}

	for _, task := range tasks {
		wg.Add(1)
		go func(t string) {
			defer wg.Done()
			fmt.Printf("Task %s started\n", t)
			time.Sleep(time.Second)
			fmt.Printf("Task %s completed\n", t)
		}(task) // pass task as an argument to avoid closure issues
	}

	wg.Wait() // wait for all goroutines to finish
	fmt.Println("All tasks completed")
}
