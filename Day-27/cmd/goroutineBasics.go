package main

import (
	"fmt"
	"time"
)

/*
Goroutines are lightweight threads managed by the Go runtime.
They allow you to execute code concurrently, and they’re created with the go keyword.

printMessage is called both in the main goroutine and in a new goroutine.
The main program will exit when the main function completes, so the goroutine
may not finish all iterations unless the main function waits.
*/

func printMessage(msg string) {
	for i := 0; i < 5; i++ {
		// print message with index number
		fmt.Println(msg, i)
		// wait for 500 milliseconds to mimic waiting / processing
		time.Sleep(500 * time.Millisecond)
	}
}

func goroutineBasics() {
	fmt.Println("\nGoroutine Basics")

	// runs in new go routine
	go printMessage("Hello world, from goroutine!")

	// runs in goroutineBasics function
	printMessage("Hello world, from goroutineBasics function!")

}
