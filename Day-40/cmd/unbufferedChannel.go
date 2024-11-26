package main

import "fmt"

func unbufferedChannels() {
	fmt.Println("\nUnbuffered Channel Blocking")

	info := `
Unbuffered channels block the sender until a receiver is ready, and vice versa.
	`
	fmt.Println(info)

	ch := make(chan int)

	go func() {
		ch <- 42 // this will block the receiver until it is ready
	}()

	// uncommenting the line below resolves the deadlock
	// fmt.Println(<-ch)
}
