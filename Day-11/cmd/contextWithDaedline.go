package main

import (
	"context"
	"fmt"
	"time"
)

/*
With context.WithDeadline, you set a specific end time for the context,
which is useful if you need operations to finish by a particular time.

context.WithDeadline allows you to specify an exact deadline (using time.Now().Add(duration)).
When the deadline expires, the context cancels, terminating any tasks that rely on it.
*/

func shortTask(ctx context.Context) {
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Short task completed")
	case <-ctx.Done():
		fmt.Println("Short task stopped:", ctx.Err())
	}
}

func contextWithDeadline() {
	fmt.Println("\nContext With Deadline")

	deadline := time.Now().Add(1 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	go shortTask(ctx)

	time.Sleep(3 * time.Second)
	fmt.Println("Main function completed")
}
