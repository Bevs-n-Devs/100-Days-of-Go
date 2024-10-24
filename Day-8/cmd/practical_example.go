package main

import (
	"fmt"
	"time"
)

/*
Here's a more complete example demonstrating Goroutines, channels,
and synchronization to simulate multiple workers performing tasks concurrently.

Three chores are created, each waiting for jobs from the jobs channel.
The main Goroutine sends five jobs to the workers and collects their results from the results channel.
*/

// chores function to simulate task processing
func chores(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second) //simulate chore by sleeping
		fmt.Printf("worker %d finished job %d\n", id, job)
		results <- job * 2 // return the results
	}
}

func practicalExercises() {
	fmt.Println("\nGoroutines, Channels and Synchronization Example")
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// start 3 chores Goroutines
	for c := 1; c <= 3; c++ {
		go chores(c, jobs, results)
	}

	// send jobs to  the chores
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // no more jobs to send

	// collect results
	for a := 1; a <= numJobs; a++ {
		result := <-results
		fmt.Printf("Result of job: %d\n", result)
	}
}
