package main

import (
	"fmt"
)

func worker(ch chan int) {
	for val := range ch {
		fmt.Println(val)
	}
}

func misconfiguredWorkerPools() {
	fmt.Println("\nMisconfigured Worker Pools")

	intro := `
Worker pools with limited channels or lack of synchronization between producers and consumers can result in deadlocks.
	`
	fmt.Println(intro)

	ch := make(chan int)
	go worker(ch)

	ch <- 1
	ch <- 2
	close(ch) // this ends the channel, causing the loop in the worker to stop

	// uncomment below to see deadlock
	// ch <- 3 // deadlock because trying to send data to a closed channel

}
