package main

import "fmt"

func closingChannels() {
	fmt.Println("\nClosing Channels")

	intro := `
A channel can be closed to signal that no more values will be sent on it.

	close(ch)
`
	fmt.Println(intro)

	ch := make(chan int, 2)

	ch <- 1
	ch <- 2
	close(ch)

	for val := range ch {
		fmt.Println(val)
	}

	explanation := `
Explanation:
- The for loop iterates over the channel until it is closed.
- After a channel is closed, it can still be read, but sending to it will cause a panic.
`
	fmt.Println(explanation)
}
