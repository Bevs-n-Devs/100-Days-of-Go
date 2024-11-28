package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: Stopping due to %v\n", id, ctx.Err())
			return
		default:
			fmt.Printf("Worker %d: Working...\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func withCancelContext() {
	fmt.Println("\nPractical Examples: Using context.WithCancel")

	// create a context with a cancel function
	ctx, cancel := context.WithCancel(context.Background())

	// start 2 workers
	go worker(ctx, 1)
	go worker(ctx, 2)

	time.Sleep(2 * time.Second)
	cancel() // signal cancellation to all workers

	time.Sleep(1 * time.Second) // wait for workers to finish
}

/* output:
Worker 1 running...
Worker 2 running...
Worker 1 running...
Worker 2 running...
Worker 1 stopping: context canceled
Worker 2 stopping: context canceled
*/
