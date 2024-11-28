package main

import "fmt"

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\n	Advanced Lesson on Contexts in Go")

	intro := `
As a mid-level developer, understanding and mastering context in Go is crucial, especially when working on applications involving timeouts, 
cancellation, or deadlines (e.g., web servers, microservices, database queries). The context package provides a way to pass deadlines, 
cancel signals, and request-scoped data across APIs and goroutines.
	`
	fmt.Println(intro)

	whatIsContext := `
	What is a Context in Go?

1. A context.Context carries deadlines, cancellation signals, and request-scoped values across function calls.
2. It is essential for managing the lifecycle of a request or a process, ensuring that:
	- Long-running operations don't block indefinitely.
	- Resources are properly cleaned up when requests are cancelled.
	`
	fmt.Println(whatIsContext)

	contextUseCases := `
	Context Use Cases:

1. Timeout Management: Stop operations if they exceed a specific time limit.
2. Cancellation Propagation: Propagate cancellation signals to goroutines to clean up resources.
3. Pass Request-Scoped Data: Attach metadata like user IDs or tracing information.
	`
	fmt.Println(contextUseCases)

	contextFunctions := `
	Context Functions:

1. context.Background(): Root context, typically used at the top level of your application.
2. context.TODO(): Placeholder context when you’re unsure about what to use.
3. context.WithCancel(parent): Derives a cancellable context from the parent.
4. context.WithTimeout(parent, timeout): Adds a timeout to the parent context.
5. context.WithDeadline(parent, deadline): Similar to timeout but uses a fixed point in time.
6. context.WithValue(parent, key, value): Passes request-scoped values down the function call chain.
	`
	fmt.Println(contextFunctions)

	withTimeoutContext()

	withCancelContext()

	withValueContext()

	CommonPitfalls := `
	Common Pitfalls:

Context Leaks: Forgetting to 1. cancel a context can lead to resource leaks.
	- Always defer cancel when creating contexts with WithCancel, WithTimeout, or WithDeadline.
2.Using context.WithValue Excessively: Overusing it can make code hard to understand.
	- Instead, pass strongly-typed arguments to functions.
3. Blocking Calls in Goroutines: If a goroutine blocks indefinitely waiting for a signal, it could cause deadlocks.
- 	Always select on ctx.Done() in your goroutines
	`
	fmt.Println(CommonPitfalls)

	bestPractices := `
	Best Practices:

1. Pass Context as the First Argument:
	- Ensure consistency by always passing ctx as the first argument in functions.

	func doSomething(ctx context.Context, param string) {
    // Implementation
	}

2. Set Reaonable Timeouts:
	- Avoid indefinite blocking operations. Use WithTimeout or WithDeadline to handle slow requests gracefully.
3. Always Check ctx.Done():
- 	In long-running loops or goroutines, periodically check ctx.Done() to handle cancellations.
4 Avoid Using context.Background Deep in the Call Chain:
	- Use context.TODO if unsure, but generally prefer passing an appropriate parent context.
6. Clean Up Resources:
	- When ctx.Done() triggers, ensure cleanup routines are executed properly.
	`
	fmt.Println(bestPractices)

	advancedConcepts := `
	Advanced Concepts:

1. Chaining Contexts:
	- You can derive multiple contexts from a parent, each with its own cancellation or timeout.

	parentCtx := context.Background()
	childCtx, cancel := context.WithTimeout(parentCtx, 1*time.Second)
	defer cancel()

2. Context in Wen Servers:
	- HTTP servers often pass a context.Context for request-scoped operations, such as database queries or external API calls.

	func handler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    // Use ctx for downstream operations
	}

3. Mock or simulate context.Context during testing to validate timeout behavior or cancellation logic.
	`
	fmt.Println(advancedConcepts)

	conclusion := `
	Final Notes 

The context package is one of Go’s most powerful tools for managing concurrency, ensuring clean shutdowns, 
and handling resource lifecycle in modern applications. Mastering context prepares you for writing scalable, 
maintainable, and resilient Go applications.	
	`
	fmt.Println(conclusion)
}
