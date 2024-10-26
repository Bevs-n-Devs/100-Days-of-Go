package main

import (
	"fmt"
	"sync"
)

/*
In scenarios where multiple Goroutines need to read a shared resource
but only a few need to write to it, you can use an RWMutex (read-write mutex).
This allows for multiple readers at the same time but only one writer at a time.

RWMutex allows multiple Goroutines to read a value concurrently but only allows one to write.
RLock and RUnlock are used for read-only operations.
Lock and Unlock are used for writing operations.
*/

var (
	counter3 int
	rwMutex  sync.RWMutex
)

func read(wg *sync.WaitGroup) {
	defer wg.Done()

	// lock the RWMutex for reading
	rwMutex.RLock()
	fmt.Println("Read counter:", counter3)
	rwMutex.RUnlock()
}

func write(wg *sync.WaitGroup) {
	defer wg.Done()

	// lock the RWMutex for writing
	rwMutex.Lock()
	counter3++
	fmt.Println("Incremented counter to:", counter3)
	rwMutex.Unlock()
}

func readWriteMutex() {
	fmt.Println("\nRead-Write Mutex")
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go write(&wg)

		wg.Add(1)
		go read(&wg)
	}

	wg.Wait()
}
