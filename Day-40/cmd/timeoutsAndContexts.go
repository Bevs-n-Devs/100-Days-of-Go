package main

import (
	"fmt"
	"time"
)

func timeoutsAndContexts() {
	fmt.Println("\nTimeouts & Contexts")

	intro := `
Use timeouts to ensure goroutines do not block indefinitely.
	`
	fmt.Println(intro)

	ch := make(chan int)

	select {
	case val := <-ch:
		fmt.Println(val)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout reached")
	}
}
