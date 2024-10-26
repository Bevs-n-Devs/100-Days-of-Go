package main

import (
	"fmt"
	"sync"
)

/*
To avoid forgetting to unlock a mutex (which can lead to deadlocks),
you can use defer to ensure that the mutex is always unlocked when the function exits.

defer mutex3.Unlock() ensures that the lock is released even if the
function exits early or encounters an error. This reduces the chances of deadlocks.
*/

var (
	counter2 int
	mutex3   sync.Mutex
)

func safeIncrement(wg *sync.WaitGroup) {
	defer wg.Done()

	// lock the critcal section
	mutex3.Lock()
	defer mutex3.Unlock() // make sure the mutex is unlocked when the function finishes

	counter2++
}

func solvingDeadlocks() {
	fmt.Println("\n\nSolving Deadlocks")

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go safeIncrement(&wg)
	}

	wg.Wait()
	fmt.Println("Final counter:", counter2)
}
