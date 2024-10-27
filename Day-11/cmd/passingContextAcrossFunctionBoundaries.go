package main

import (
	"context"
	"fmt"
	"time"
)

/*
Contexts are especially useful when passed across multiple functions or
Goroutines, allowing cancellation or timeout behavior to propagate across an entire task.

In this example, parentFunction calls childFunction with the same context.
When the context is canceled, all functions in the chain stop execution.
*/
func parentFunction(ctx context.Context) {
	fmt.Println("Parent function started")
	childFunction(ctx)
}

func childFunction(ctx context.Context) {
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Child function completed")
	case <-ctx.Done():
		fmt.Println("Child function cancelled:", ctx.Err())
	}
}

func passingContextAcrossFunctionBoundaries() {
	fmt.Println("\nPassing Context Across Function Boundaries")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	parentFunction(ctx)

	time.Sleep(3 * time.Second)
	fmt.Println("Main function completed")
}
