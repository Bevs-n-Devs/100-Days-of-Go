package main

import (
	"fmt"
	"sync"
)

var (
	data   = 0
	rwLock sync.RWMutex
)

func reader(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	rwLock.RLock() // aquire read lock
	defer rwLock.RUnlock()
	fmt.Printf("Reader %d read data: %d\n", id, data)
}

func writer(wg *sync.WaitGroup, value int) {
	defer wg.Done()
	rwLock.Lock() // aquire write lock
	defer rwLock.Unlock()
	data = value // modify data
	fmt.Printf("Writer updated data to: %d\n", data)
}

func advacned() {
	fmt.Println("\n	Advanced: RWMutex")

	intro := `
Go provides sync.RWMutex, a specialized mutex for scenarios where 
multiple readers can access data simultaneously, but only one writer can modify it.
	`
	fmt.Println(intro)

	var wg sync.WaitGroup

	wg.Add(3)
	go reader(&wg, 1)
	go writer(&wg, 42)
	go reader(&wg, 2)

	wg.Wait()
}
