package main

import (
	"fmt"
	"sync"
)

func worker2(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
	}
}

func waitGroupsAndChannels() {
	fmt.Println("\nUsing WaitGroups and Channels")

	var wg sync.WaitGroup
	jobs := make(chan int, 10)

	// start workers
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker2(w, jobs, &wg)
	}

	// send jobs
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs) // close channel to signal no more jobs

	wg.Wait()
	fmt.Println("All jobs completed")
}
