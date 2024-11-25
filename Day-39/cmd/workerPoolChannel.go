package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		results <- job * 2 // return the results
	}
}

func workerPoolChannel() {
	fmt.Println("\nWorker Pool Channel")

	info := `
Worker Pools Channels are often used to manage a pool of worker goroutines.
	`
	fmt.Println(info)

	jobs := make(chan int, 5)
	results := make(chan int, 5)
	var wg sync.WaitGroup

	// launch workers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// submit jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// we use the wait group to wait for all the workers to finish their jobs
	// once all the workers are done, we close the channel of results to signal that there are no more results coming in
	// this is important because we use a range loop to retrieve all the results, and the range loop will block until the channel is closed
	wg.Wait()
	close(results)

	// get results
	for result := range results {
		fmt.Println("Result:", result)
	}

}
