package main

import "fmt"

func bestPractices() {
	fmt.Println("\n	Best Practices")

	bestPractices := `
1. Minimize the Scope of Locks: Hold the lock for the shortest possible time to reduce contention:

	mu.Lock()
	value := sharedVar
	mu.Unlock()
	fmt.Println(value) // Outside critical section

2. Avoid Deadlocks: Deadlocks occur when two or more goroutines wait indefinitely for each other to release locks. Strategies to prevent deadlocks:
	- Always acquire multiple locks in the same order.
	- Use timeouts or try-lock mechanisms (e.g., sync/atomic for simple tasks).
3. Prefer Higher-Level Constructs: If your use case allows, prefer channels or other synchronization tools (e.g., sync.Map, sync.RWMutex).
	`
	fmt.Println(bestPractices)
}
