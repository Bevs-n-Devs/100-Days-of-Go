package main

import (
	"fmt"
	"sync"
)

var counter = 0

func increment(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		counter++ // data race: multiple goroutines modify the same counter variable simultaneously
	}

}

func whyUseMutex() {
	fmt.Println("\n	Why Use A Mutex?")

	intro := `
Without a mutex, multiple goroutines accessing shared data concurrently may lead to data races or corrupted states.
	`
	fmt.Println(intro)

	var wg sync.WaitGroup
	wg.Add(2)

	go increment(&wg)
	go increment(&wg)

	wg.Wait()
	fmt.Println("Final Counter Value:", counter) // likely to be incorrect due to race conditions
}
