package main

import "fmt"

/*
Channels can be closed when no more values will be sent to them.
Receivers can check if a channel is closed by using the comma ok idiom.

The sender sends five values to the channel and then closes it.
The receiver can iterate over the channel using the for val := range ch syntax,
which automatically stops when the channel is closed.
*/

func closingChannels() {
	fmt.Println("\nClosing Channels")

	ch := make(chan int)

	// start a Goroutine to send values
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch) // close the channel after sending all values
	}()

	// receive values until the channel is closed
	for val := range ch {
		fmt.Println(val)
	}
}
