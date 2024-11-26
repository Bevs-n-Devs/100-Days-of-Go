package main

import (
	"fmt"
	"sync"
)

func improperUseofMutexes() {
	fmt.Println("\nImproper Use of Mutexes")

	intro := `
When locks are not released properly, deadlocks can occur.
	`
	fmt.Println(intro)

	var mu sync.Mutex

	mu.Lock()
	// uncomment below to see deadlock
	// mu.Lock() // Deadlock: the same goroutine cannot lock twice without unlocking
	fmt.Println("The same goroutine cannot lock twice without unlocking")
}
