package main

import "fmt"

/*
 */

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nDeep Dive into Contexts")

	intro := `
What we will cover:

1. Understand Context Basics: How context works and when to use it.
2. Creating and Using Contexts: How to create, pass, and utilize context.
3. Context Types and Scenarios: context.Background(), context.WithCancel, context.WithDeadline, and context.WithTimeout.
4. Best Practices for using contexts in concurrent programs.`

	whatIsContext := `
What is a Context in Go?
 	
A Context provides a way to:

- Pass request-scoped values and metadata.
- Signal a timeout or cancellation to long-running processes.
-Control the lifespan of goroutines, ensuring they don't run indefinitely.
`
	usedIn := `
Contexts are often used in:

- HTTP servers to manage request lifecycles.
- Database queries to manage timeouts.
- Background tasks to handle cancellation if a task is no longer needed.
`

	usingContexts := `
Using Contexts in Go:

Contexts are designed to be passed from the main routine to goroutines and functions that need them. Here's an overview of some functions in the context package:

1. context.Background(): Creates an empty context, often used as a base.
2. context.WithCancel(ctx): Creates a context that can be canceled manually.
3. context.WithTimeout(ctx, duration): Creates a context that will automatically cancel after a specified duration.
4. context.WithDeadline(ctx, time): Similar to WithTimeout, but instead of a duration, you provide a specific point in time when the context will cancel.`

	fmt.Println(intro)
	fmt.Println(whatIsContext)
	fmt.Println(usedIn)
	fmt.Println(usingContexts)

	contextExample()

	contextInAPIs()
}

/*
BEST PRACTICES

1. Pass Context as the First Parameter:
- Always pass context.Context as the first parameter to functions that require it.
- This is the convention in Go and makes it clear which functions are context-aware.

2. Avoid Contexts in Structs:
- Contexts are designed to be short-lived, so it’s not recommended to store them in structs or global variables.
  Instead, pass them directly to functions or goroutines that need them.

3. Use context.TODO() for Incomplete Contexts:
- If you know you’ll need a context but aren’t sure what type yet, use context.TODO().
  This is helpful when planning for context implementation later in development.

4. Always Call Cancel:
- If you use context.WithCancel, context.WithTimeout, or context.WithDeadline, always call the cancel function when done, typically using defer.
  This ensures that resources tied to the context are released properly.

5. Do Not Pass nil Contexts:
- Always use a real context (context.Background() or context.TODO()) as the base instead of passing nil to functions.
*/
