package main

import (
	"fmt"
	"time"
)

/*
Three workers are started as Goroutines.
Each worker waits for 2 seconds and then sends its ID to the channel when it finishes.

The main function receives messages from the channel and prints which worker has completed its work.

GO TIP:
Do NOT communicate by sharing memory; intead share memory by COMMUNICATING!
*/

// function to simulate work in a Goroutine
func worker(id int, ch chan int) {
	time.Sleep(2 * time.Second)
	fmt.Printf("Worker %d finished work\n", id)
	ch <- id // send the worker's ID to the channel
}

func goroutinesAndChannels() {
	fmt.Println("\nGoroutines and Channels: Working together")

	// create channel to communicate between Goroutines
	ch := make(chan int)

	// start 3x workers (Goroutines)
	for i := 1; i <= 3; i++ {
		go worker(i, ch)
	}

	// wait fro each Goroutine to finish & get message from channel
	for i := 1; i <= 3; i++ {
		workerID := <-ch // get workerID from the channel
		fmt.Printf("Received completion signal from the worker %d\n", workerID)
	}

	fmt.Println("All workers finished!")
}
