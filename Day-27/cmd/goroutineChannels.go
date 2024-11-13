package main

import (
	"fmt"
	"time"
)

/*
Channels allow goroutines to communicate with each other, passing messages or data safely between them.

The channel ch is used to communicate between main and worker goroutines.
Values are sent from worker and received in main, enabling synchronization.
*/

func worker(id int, ch chan<- int) {
	fmt.Printf("Worker %d started\n", id)
	// send value to the channel
	ch <- id * 2
}

func goroutineChannel() {
	fmt.Println("\nGoroutines with (unbuffered) Channels")

	// create a new (unbuffered) channel
	ch := make(chan int)

	go worker(1, ch)
	go worker(2, ch)

	// recieve values from channel
	result1 := <-ch
	result2 := <-ch

	fmt.Println("Results:", result1, result2)

	fmt.Println("\nGoroutines with (buffered) Channels")
	// create new bufferd channel
	ch2 := make(chan int, 3) // max capacity of 3

	// send 3 values to channel 2
	fmt.Println("Sending 10 to channel 2")
	ch2 <- 10
	fmt.Println("Sending 20 to channel 2")
	ch2 <- 20
	fmt.Println("Sending 30 to channel 3")
	ch2 <- 30

	// this next send will block becasue the channel is full
	go func() {
		fmt.Println("Attempting to send 40 to channel 2")
		ch2 <- 40 // this will block until there is a reciever
		fmt.Println("Sent 40 to channel 2")
	}()

	// give some delay to print the statements
	time.Sleep(2 * time.Second)

	// start recieving from channel - this will unblock goroutine trying to send 40 as the capacity will no longer be full
	fmt.Println("Receiving from channel 2:", <-ch2) // gets 10
	fmt.Println("Receiving from channel 2:", <-ch2) // gets 20
	fmt.Println("Receiving from channel 2:", <-ch2) // gets 30

	// add delay to show unblocking sequence
	time.Sleep(1 * time.Second)

	// Recieving one more allows the goroutine to send the bolocked 40
	fmt.Println("Receiving from channel 2:", <-ch2) // gets 40
}

/*
KEY POINTS FOR UNBUFFERED CHANNELS
1. Synchronization tool: Used for coordinating two goroutines.
2. Synchronous send/receive: Both the sender and receiver must be ready simultaneously.
3. No storage: Without any buffer, unbuffered channels can only hold data momentarily during handoff.

WHEN TO AVOID UNBUFFERED CHANNELS
1. If High Throughput is Needed: In scenarios where data production is much faster than consumption,
unbuffered channels can slow down the program due to their synchronous nature. Here, a buffered channel or worker pool may be more efficient.

2. Non-Blocking Communication: If the sender does not need to wait for a receiver, such as when a
producer can generate multiple pieces of data without an immediate consumer, a buffered channel is better suited.

WHEN TO USE BUFFERED CHANNELS
1. Rate Limiting: You can control how many tasks or events are processed at once.
2. Decoupling Goroutines: When a producer and a consumer do not need to be perfectly synchronized, a buffered channel provides temporary storage.
3. Minimizing Blocking: Buffering lets the producer keep running up to a certain limit before it blocks.
*/
