package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
In some cases, you might want to avoid the complexity of mutexes.
Go’s sync/atomic package provides low-level atomic memory primitives that can be used to
perform atomic operations like incrementing a counter without using a mutex.

atomic.AddInt64 is used to increment the counter atomically, which means the operation is safe to run concurrently.
Atomic operations can be faster and simpler than mutexes but are typically limited to basic operations (increment, decrement, etc.).
*/

var counter4 int64

func incrementCounter(wg *sync.WaitGroup) {
	defer wg.Done()

	// perform atomic increment
	atomic.AddInt64(&counter4, 1)
}

func atomicOperations() {
	fmt.Println("\nAtomic Operations")

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go incrementCounter(&wg)
	}

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
