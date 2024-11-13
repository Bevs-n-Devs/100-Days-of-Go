package concurrencypatterns

import (
	"fmt"
	"sync"
)

/*
Worker pools allow for concurrent task processing, ideal for limiting the number of concurrent tasks (e.g., limiting I/O or CPU usage).

We create 3 worker goroutines, and each worker processes tasks from the tasks channel. The main function waits for all workers to finish using a sync.WaitGroup.
*/

func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // makes sure the goroutine completes safely
	for task := range tasks {
		fmt.Printf("Worker %d processing task %d\n", id, task)
		// multiply the task by 2 and save to results channel
		results <- task * 2
	}
}

// funcion needs to start with capital letter to be used globally throughout project directory
func WorkerPools() {
	fmt.Println("\nConcurrency Patterns: Worker Pool")

	// create channels and WaitGroup
	tasks := make(chan int, 10)
	results := make(chan int, 10)
	var wg sync.WaitGroup

	// launch workers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, tasks, results, &wg)
	}

	// send tasks
	for i := 1; i <= 5; i++ {
		tasks <- i
	}
	close(tasks)

	// collect results
	go func() {
		wg.Wait()
		close(results)
	}()
	for result := range results {
		fmt.Println("Result:", result)
	}
}
