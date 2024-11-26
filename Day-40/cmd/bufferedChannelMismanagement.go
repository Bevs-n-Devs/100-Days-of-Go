package main

import "fmt"

func bufferedChannelMismanagement() {
	fmt.Println("\nBuffered Channel Mismanagement")

	intro := `
Buffered channels block only when the buffer is full (for sends) or empty (for receives). Mismanaging the buffer size can lead to deadlocks.
	`
	fmt.Println(intro)

	ch := make(chan int, 2) // create a buffered channel with a capacity of 2
	ch <- 1
	ch <- 2
	// uncomment below to see deadlock
	// ch <- 3 // Deadlock: the buffer is full and no goroutine is recieving data
	fmt.Println(<-ch)
}
