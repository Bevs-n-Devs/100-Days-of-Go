package main

import (
	"fmt"
	"sync"
)

/*
A deadlock occurs when two or more Goroutines wait indefinitely for a
condition that will never be satisfied, often caused by improper locking and unlocking of shared resources.

In this example, the program will get stuck because it tries to lock the same mutex twice without unlocking it first.
Be careful when locking to avoid deadlocks by ensuring that each lock is properly paired with an unlock.
*/

func deadlock() {
	fmt.Println("\nDeadlocks")

	var mutex sync.Mutex

	mutex.Lock() // lock the mutex

	fmt.Println("First lock acquired")

	// attempt ot lock again without unlocking (this will cause a deadlock)
	mutex.Lock()

	fmt.Println("Second lock acquired")
}
