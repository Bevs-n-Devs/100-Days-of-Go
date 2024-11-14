package main

import (
	"context"
	"fmt"
	"time"
)

/*
Context with Timeout:
- context.WithTimeout(context.Background(), 2*time.Second): Creates a context that will automatically cancel after 2 seconds.
- This timeout is useful for tasks that should complete within a specific period. In this example, task 1 stops after 2 seconds when the context expires.

Context with Cancel:
- context.WithCancel(context.Background()): Creates a context with no timeout but with a cancel function.
- After 2 seconds, cancel() is called manu
*/

// simulate a task that checks for context cancellation
func doTask(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Task %d: Stopping due to %v\n", id, ctx.Err())
			return
		default:
			fmt.Printf("Task %d: Running...\n", id)
			time.Sleep(500 * time.Millisecond) // simulate work
		}
	}
}

func contextExample() {
	fmt.Println("\nContext Example")

	// example using context with timeout
	fmt.Println("Example: Context with Timeout")
	cxtTimeout, cancelTimeout := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelTimeout()

	go doTask(cxtTimeout, 1) // start a goroutine with a timeout context

	time.Sleep(3 * time.Second) // let it run for 3 seconds
	fmt.Println("Main function ends")

	// example using context with manual cancel
	fmt.Println("\nExample: Context with cancel")
	ctxCancel, cancel := context.WithCancel(context.Background())
	go doTask(ctxCancel, 2) // start a goroutine with a context you can cancel

	time.Sleep(2 * time.Second) // run a task for 2 seconds
	cancel()                    // manually cancel the context
	time.Sleep(1 * time.Second) // let the goroutine finish
}
