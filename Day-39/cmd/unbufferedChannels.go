package main

import "fmt"

func unbufferedChannels() {
	fmt.Println("\nUnbuffered Channels")

	info := `An unbuffered channel has no capacity to store values. 
Every send (ch <- value) operation will block until another goroutine performs a corresponding receive (<-ch) operation.
	`
	fmt.Println(info)

	ch := make(chan string) // unbuffered channel

	// goroutine to send data to the channel
	go func() {
		ch <- "Hello from the goroutine"
	}()

	// receive data from the channel
	msg := <-ch
	fmt.Println(msg)

	explanation := `
Explanation:
- The ch <- "Hello from goroutine" blocks until the main function retrieves the message using <-ch.
- This synchronization ensures that both sender and receiver are coordinated.
	`
	fmt.Println(explanation)
}
