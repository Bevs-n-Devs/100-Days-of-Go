package main

import "fmt"

func fanOutWorker(id int, ch <-chan int) {
	for job := range ch {
		fmt.Printf("Worker %d recieved job %d\n", id, job)
	}
}

func fanOUt() {
	fmt.Println("\nFan Out")

	intro := `
Fan-Out: Multiple goroutines read from the same channel.	
	`
	fmt.Println(intro)

	ch := make(chan int)

	// start 3x workers
	for i := 1; i <= 3; i++ {
		go fanOutWorker(i, ch)
	}

	// send jobs to the channel
	for job := 1; job <= 5; job++ {
		ch <- job
	}
	close(ch)
}
