package concurrencypatterns

import (
	"fmt"
	"sync"
)

/*
Fan-Out means having multiple goroutines read from the same channel, while Fan-In means multiple goroutines writing to a single channel.
*/

func FanInFanOut() {
	fmt.Println("\nConcurrency Patterns: Fan-In / Fan-Out")

	// make channels & WaitGroup
	tasks := make(chan int, 5)
	results := make(chan int, 5)
	var wg sync.WaitGroup

	// Fan-Out: start multiple workers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, tasks, results, &wg)
	}

	// send tasks
	for i := 1; i <= 5; i++ {
		tasks <- i
	}
	close(tasks)

	// Fan-In: collect all results
	go func() {
		wg.Wait()
		close(results)
	}()
	for result := range results {
		fmt.Println("Result:", result)
	}
}
