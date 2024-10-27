package main

import (
	"context"
	"fmt"
)

/*
Contexts can also carry request-scoped values, allowing you to store and retrieve data across function boundaries.
Note that contexts should be used for values that control the execution of a task, like authentication tokens,
request IDs, or other short-lived data. Avoid using contexts for long-term data storage.

context.WithValue adds a key-value pair to the context, making it available to any functions that share the same context.
In this example, userID is added to the context and accessed within processRequest.

Note: Use constants for keys when setting context values to prevent collisions.

	Avoid using context for data not directly related to controlling the execution.
*/
func processRequest2(ctx context.Context) {
	userID := ctx.Value("userID")
	if userID == nil {
		fmt.Println("No user ID found in context")
		return
	}
	fmt.Println("Processing request for user ID:", userID)
}

func usingContextWithValues() {
	fmt.Println("\nUsing Context With Values")

	ctx := context.WithValue(context.Background(), "userID", 42)

	processRequest2(ctx)
}
