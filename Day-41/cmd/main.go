package main

import "fmt"

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nDeep Dive Into Mutexes")

	intro := `
	What is a Mutex?

A Mutex is a synchronization tool that allows only one goroutine to access a critical section of code at a time. 
Mutexes are used to protect shared resources (e.g., variables, data structures) from concurrent modification, which can cause data races.
	- Lock: Acquires the mutex. If another goroutine already has the lock, the calling goroutine blocks until the lock is released.
	- Unlock: Releases the mutex, allowing other goroutines to acquire it.
	`
	fmt.Println(intro)

	whyUseMutex()

	usingMutexCorrectly()

	howMutexWork := `
	How Mutexes Work?

Internally, a mutex maintains a lock state and a queue of waiting goroutines:

1. Locking:
	- If the mutex is free, the calling goroutine acquires the lock.
	- If it’s already locked, the calling goroutine blocks until it becomes available.
2. Unlocking:
	- Releases the lock and wakes up one of the waiting goroutines.
	`
	fmt.Println(howMutexWork)

	commonMistakes()

	bestPractices()

	advacned()

	mutexVsChannels := `
	Mutext vs. Channels

Both sync.Mutex and channels can be used for synchronization, but they have different trade-offs:
	- Use mutexes for protecting shared memory.
	- Use channels for communicating between goroutines without shared memory.
	`
	fmt.Println(mutexVsChannels)
}
