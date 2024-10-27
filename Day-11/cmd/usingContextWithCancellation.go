package main

import (
	"context"
	"fmt"
	"time"
)

/*
When you need more control, context.WithCancel creates a context that you can cancel manually.

The worker function checks ctx.Done() in a loop, which allows it to terminate gracefully when canceled.
Calling cancel() triggers the context cancellation, ending the worker function.
*/
func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker stopped:", ctx.Err())
			return
		default:
			fmt.Println("Worker is running...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func usingContextWithCancellation() {
	fmt.Println("\nUsing Context With Cancellation")

	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx)

	time.Sleep(2 * time.Second)
	fmt.Println("Cancelling the context...")
	cancel() // manully cancel the context

	time.Sleep(1 * time.Second)
	fmt.Println("Main function completed")
}
