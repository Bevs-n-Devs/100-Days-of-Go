package main

import (
	"fmt"
	"sync"
	"time"
)

// worker function to simulate work in a Goroutine
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // defer is used to ensure the function is executed even if an error occurs
	fmt.Printf("Worker %d started\n", id)

	// simulate work
	time.Sleep(time.Second)
	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nDay 38 of 100 Days of Go: Waitgroups")

	intro := `
In Go, concurrency is a fundamental aspect, and managing multiple goroutines effectively is crucial for building robust applications. 
Wait Groups (sync.WaitGroup) are one of the essential tools provided by Go to synchronize goroutines.
	`
	fmt.Println(intro)

	whatIsWaitGroup := `
	What is a Wait Group?

A Wait Group is used to wait for a collection of goroutines to finish executing. 
It provides a way to track multiple concurrent operations and ensures the main 
program or controlling goroutine does not exit prematurely before all other goroutines complete.

Wait Groups are part of the sync package and work with three main methods:
	1. Add(delta int): Adds a counter indicating the number of goroutines to wait for.
	2. Done(): Decrements the counter when a goroutine completes.
	3. Wait(): Blocks until the counter becomes zero, meaning all goroutines have finished.	
	`
	fmt.Println(whatIsWaitGroup)

	// basic waitgroup example
	fmt.Println("\nBasic Wait Group Example")
	var wg sync.WaitGroup

	// iterate over 3 workers
	for i := 1; i <= 3; i++ {
		wg.Add(1)         // increment the counter for each goroutine
		go worker(i, &wg) // start the goroutine
	}

	// each worker will call wg.Done() when it finishes
	// the Wait() method will block until all goroutines have finished

	wg.Wait() // wait for all goroutines to finish
	fmt.Println("All goroutines have finished.")

	howWaitGroupWorks := `
	How Wait Groups Work

1. Start Tracking Tasks with Add
	- Increment the counter with the number of goroutines to wait for.
2. Run Goroutines
	- Spawn the goroutines that need synchronization.
3. Mark Completion with Done
	- Each goroutine calls Done() to signal completion.
4. Wait for All Tasks with Wait
	- Blocks the main program or calling goroutine until the counter becomes zero.
	`
	fmt.Println(howWaitGroupWorks)

	waitGroupRules := `
	Important Rules for Using Wait Groups

1. Always Pair Add and Goroutines
	- Ensure you call wg.Add(1) before starting the goroutine. Otherwise, there's a risk of missing a Done() call.
2. Avoid Calling Wait Inside a Goroutine
	- Wait blocks until the counter becomes zero. Calling it inside a goroutine can lead to deadlocks.
3. Use defer wg.Done()
	- Placing defer wg.Done() at the start of the goroutine ensures the counter is decremented even if the goroutine exits prematurely due to an error or panic.
	`
	fmt.Println(waitGroupRules)

	waitGroupsWithAnonymousFunctions()

	waitGroupsWithMultipleResources()

	waitGroupsAndChannels()

	bestPractices := `
	BEST PRACTICES

1. Use a Pointer for the Wait Group
	- Always pass a pointer to the sync.WaitGroup to avoid unnecessary copies and ensure the same counter is modified.
2. Defer wg.Done() Immediately
	- Place defer wg.Done() at the beginning of each goroutine to ensure it executes even if the goroutine encounters an error or panic.
3. Avoid Overusing Wait Groups
	- If you need to synchronize complex workflows, consider using channels or sync.Once alongside wait groups.
4. Don’t Reuse Wait Groups
	- A Wait Group should be used for a single operation. Reusing it can lead to undefined behavior.
	`
	fmt.Println(bestPractices)

	commonMistakes := `
	Common Mistakes

1. Missing wg.Done() Forgetting to call Done() results in the Wait method hanging indefinitely.
2. Calling Wait Before All Goroutines Are Added Calling Wait() too early might result in a race condition or goroutines being ignored.
3. Modifying a Wait Group Counter Concurrently Only one goroutine (usually the main one) should call Add.	
	`
	fmt.Println(commonMistakes)

	conclusion := `
	CONCLUSION

Wait Groups are a fundamental concurrency primitive in Go that make it easy to synchronize multiple goroutines. 
By understanding how to use Add, Done, and Wait, you can coordinate goroutines effectively, avoid race conditions, 
and ensure your program runs correctly. Combine wait groups with other concurrency tools like channels for more 
robust and flexible patterns.
	`
	fmt.Println(conclusion)

}
