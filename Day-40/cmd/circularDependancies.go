package main

import (
	"fmt"
	"sync"
)

func circularDependencies() {
	fmt.Println("\nCircular Dependancies")

	intro := `
Goroutines waiting on each other to release resources can lead to circular dependencies.
	`
	fmt.Print(intro)

	var mu1, mu2 sync.Mutex

	go func() {
		mu1.Lock()
		defer mu1.Unlock()

		mu2.Lock()
		defer mu2.Unlock()
	}()

	fmt.Println("Deadlock occurs!")
}
