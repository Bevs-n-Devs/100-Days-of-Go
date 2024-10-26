package main

import (
	"fmt"
	"sync"
)

/*
Channels are useful for communication between Goroutines. They allow for passing data between Goroutines safely without needing explicit locks.
For many cases, channels can be a simpler way to manage concurrency, but mutexes provide more direct control over shared resources.

In the mutex example, the mutex controls the access to the shared counter variable.
*/

var (
	counter5 int
	mutex4   sync.Mutex
)

func incrementNumber(wg *sync.WaitGroup) {
	defer wg.Done()

	mutex4.Lock()
	counter5++
	mutex4.Unlock()
}

func usingMutex() {
	fmt.Println("\nChannels vs Mutex: Using A Mutex")

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go incrementNumber(&wg)
	}

	wg.Wait()
	fmt.Println("Final Counter with Mutex:", counter5)
}
