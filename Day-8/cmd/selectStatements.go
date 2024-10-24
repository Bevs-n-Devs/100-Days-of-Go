package main

import (
	"fmt"
	"time"
)

/*
The select statement allows a Goroutine to wait on multiple communication
operations (channels). It’s like a switch statement but for channels.

The select statement waits for messages from both ch1 and ch2 channels.
Whichever channel sends a message first will be handled, and the other case will be ignored.
This is useful for handling multiple asynchronous operations concurrently.

*/

func selectStatements() {
	fmt.Println("\nSelect Statements with Goroutines and Channels")

	ch1 := make(chan string)
	ch2 := make(chan string)

	// start 2x Goroutines to send values to the channels
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Message from Channel 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Message from Channel 2"
	}()

	// use select to wait for either channel to send a message
	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}

	fmt.Println("Finished!")
}
