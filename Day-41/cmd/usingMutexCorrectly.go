package main

import (
	"fmt"
	"sync"
)

var (
	counter2 = 0
	mtx      sync.Mutex
)

func increment2(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		mtx.Lock()   // lock the mutext before modifying the counter variable
		counter2++   // safely increemnt the counter variable
		mtx.Unlock() // unlock the mutext
	}
}

func usingMutexCorrectly() {
	fmt.Println("\nUsing Mutex Correctly")

	intro := `
To fix the data race, wrap the shared resource access in a critical section using a sync.Mutex.
	`
	fmt.Println(intro)

	var wg sync.WaitGroup
	wg.Add(2)

	go increment2(&wg)
	go increment2(&wg)

	wg.Wait()
	fmt.Println("Final Counter Value:", counter2) // correct result
}
