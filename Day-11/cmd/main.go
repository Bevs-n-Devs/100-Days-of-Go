package main

import (
	"context"
	"fmt"
)

/*
The context package is essential for managing timeouts, cancellations, and deadlines in concurrent operations.
The context package is frequently used in web applications, microservices, and other systems where you
need control over long-running or potentially cancellable tasks.

The context package provides a way to carry deadlines, cancellation signals, and other request-scoped
values across API boundaries and between Goroutines.

Contexts are created with three main types of functions:

context.Background() - Used as the root context in applications, often in main or initialization functions.
context.WithTimeout(ctx, duration) - Sets a timeout for a context.
context.WithCancel(ctx) - Allows you to cancel an operation manually.
context.WithDeadline(ctx, time) - Sets a specific deadline for a context.

context.Background() is usually used as the starting context.
From this, you can derive other contexts with specific configurations.
*/

func main() {
	fmt.Println("Hello world, hello Yaw!")
	fmt.Println("\nContext Package - Managing Timeouts, Cancellation, and Deadlines in Go")

	ctx := context.Background() // create a root context
	fmt.Println("Root context created:", ctx)

	usingContextWithTimeout()

	usingContextWithCancellation()

	contextWithDeadline()

	passingContextAcrossFunctionBoundaries()

	usingContextWithValues()

	realWorldApplication()
}
