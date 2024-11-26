package main

import "fmt"

func commonMistakes() {
	fmt.Println("\n	Common Mistakes And Pitfalls")

	forgettingUnlock := `
Forgetting to Unlock

	mu.Lock()
	// Critical section
	// Missing mu.Unlock() causes a deadlock!

If you forget to call .Unlock(), other goroutines will block indefinately, leading to a deadlock.
The way to prevent this is to use a defer statement.

	mu.Lock()
	defer mu.Unlock()
	`
	fmt.Println(forgettingUnlock)

	doubleLocking := `
Double Locking

	mu.Lock()
	mu.Lock() // Deadlock: The same goroutine tries to acquire the lock twice.
	`
	fmt.Println(doubleLocking)

	unlockingUnlockedMutex := `
Unlocking an Unlocked Mutex

Unlocking a mutext that wasn't locked causes a runtime panic.

	mu.Unlock()  // Panic: unlock of unlocked mutex
	`
	fmt.Println(unlockingUnlockedMutex)
}
