package main

import (
	"context"
	"fmt"
	"time"
)

func queryDatabase(ctx context.Context) {
	select {
	case <-time.After(3 * time.Second): // simulate a long database query
		fmt.Println("Query successful!")
	case <-ctx.Done(): // context cancellation or timeout
		fmt.Println("Query cancelled:", ctx.Err())
	}
}

func withTimeoutContext() {
	fmt.Println("\nPratical Eaxmples: Using context.WithTimeOut")

	// create a context with a 2 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // ensure resources are cleaned up

	queryDatabase(ctx)
}

/* output:
Query cancelled: context deadline exceeded
*/
