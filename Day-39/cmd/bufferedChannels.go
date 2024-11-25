package main

import "fmt"

func bufferedChannels() {
	fmt.Println("\nBuffered Channels")

	intro := `A buffered channel has a fixed capacity to store values, allowing send operations to proceed without blocking until the buffer is full.
To create a buffered channel, you can use the make() function with the second argument specifying the capacity.

	ch := make(chan int, 3) // Buffered channel with capacity 3
	`
	fmt.Println(intro)

	ch := make(chan int, 3) // Buffered channel with capacity 3

	ch <- 1 // non-blocking send
	ch <- 2 // non-blocking send

	fmt.Println(<-ch) // gets the first value from the channel
	fmt.Println(<-ch) // gets the second value from the channel

	keyPoints := `

Key Points:
- Buffered channels have a fixed capacity to store values, allowing send operations to proceed without blocking until the buffer is full.
- The capacity of the channel is specified as the second argument when creating the channel using make().
- Send operations block if the buffer is full.
- Receive operations block if the buffer is empty.
	`
	fmt.Println(keyPoints)
}
