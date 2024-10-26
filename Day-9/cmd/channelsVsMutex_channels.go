package main

import (
	"fmt"
	"sync"
)

/*
Channels are useful for communication between Goroutines.
They allow for passing data between Goroutines safely without needing explicit locks.

In the channel example, a single channel controls access to the counter value,
and each Goroutine receives and updates the value in sequence.
*/

func incrementNumber2(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	// receive the current value, increment, and send it back
	value := <-ch
	value++
	ch <- value
}

func usingChannel() {
	fmt.Println("\nChannels vs Mutex: Using Channels")

	var wg sync.WaitGroup
	ch := make(chan int, 1)
	ch <- 0 // initialize the channel with the counter value

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go incrementNumber2(ch, &wg)
	}

	wg.Wait()
	counter6 := <-ch // get the final counter value
	fmt.Println("Final Counter with Channel:", counter6)
}
