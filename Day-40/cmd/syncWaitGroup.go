package main

import (
	"fmt"
	"sync"
)

func syncWaitGroup() {
	fmt.Println("\nUsing sync.WaitGroup")

	intro := `
Use sync.WaitGroup to ensure all goroutines complete before the main function exits.
	`
	fmt.Println(intro)

	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		ch <- 42
	}()

	go func() {
		wg.Add(1)
		defer wg.Done()
		fmt.Println(<-ch)
	}()

	wg.Wait()

}
