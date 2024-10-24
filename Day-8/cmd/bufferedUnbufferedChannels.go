package main

import "fmt"

/*
Unbuffered Channels (like the ones in the previous examples) block the sender until the receiver is ready,
and block the receiver until a value is available.
Buffered Channels allow sending values without waiting for a receiver as long as the buffer isn’t full.

The channel can hold up to two values before blocking.
The values can be sent to the channel without needing an immediate receiver.
*/

func bufferdChannels() {
	fmt.Println("\nBuffered Channels")

	ch := make(chan int, 2) // create a bufferd channel with a capacity of 2

	ch <- 16
	ch <- 10

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
