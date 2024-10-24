package main

import (
	"fmt"
	"time"
)

/*
You can also start Goroutines with anonymous functions

The anonymous function starts running in a new Goroutine, allowing concurrent execution with the goroutineAnonymousFunctions function
*/

func goroutineAnonymousFunctions() {
	go func() {
		fmt.Println("\nGoroutuines and Anonymous Functions")
		fmt.Println("This is a Goroutine witrh an anonymous function")
	}()

	// wait for Goroutine to finish
	time.Sleep(1 * time.Second)
}
