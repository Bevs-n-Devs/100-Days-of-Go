package main

import (
	"context"
	"fmt"
	"time"
)

/*
A timeout context automatically cancels itself if the
work doesn’t complete within the specified time.

context.WithTimeout: Creates a context that cancels itself after 1 second.
ctx.Done(): Channel that closes when the context is canceled.
In this example, the longRunningTask function is canceled by the timeout
before it completes, and ctx.Err() gives the reason for cancellation.
*/
func longRunningTask(ctx context.Context) {
	select {
	case <-time.After(2 * time.Second): // simulate work for 2 seconds
		fmt.Println("Task completed")
	case <-ctx.Done(): // listen for the context's cancellation signal
		fmt.Println("Task cancelled:", ctx.Err())
	}
}

func usingContextWithTimeout() {
	fmt.Println("\nUsing Context With Timeout")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go longRunningTask(ctx)

	time.Sleep(3 * time.Second)
	fmt.Println("Main function completed")
}
