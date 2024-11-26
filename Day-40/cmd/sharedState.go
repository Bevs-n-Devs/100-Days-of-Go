package main

import "fmt"

func sharedState() {
	fmt.Println("\nAvoid Shared State")

	intro := `
Favour communication over shared memory. Instead of locking, use channels to coordinate goroutines.
	`
	fmt.Println(intro)

	ch := make(chan int)

	go func() {
		ch <- 42
	}()

	fmt.Println(<-ch)
}
