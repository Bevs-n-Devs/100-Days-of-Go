package main

import "fmt"

func whatIsDeadlock() {
	intro := `
	What is a Deadlock?

A deadlock occurs when:
1. Two or more goroutines are blocked, waiting for a condition that will never be met.
2. None of the goroutines can proceed because they are mutually dependent on each other to release resources.
	`
	fmt.Println(intro)

	ch := make(chan int)
	ch <- 1 // deadlock: no goroutine is recieving data from the channel
	fmt.Println(<-ch)
}

// The send operation (ch <- 1) blocks because there is no corresponding receive operation at the same time.
