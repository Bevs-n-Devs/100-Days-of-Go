package main

import (
	"fmt"
)

/*
Non-blocking operations are helpful when you want to check a
channel without blocking the Goroutine. You can use the default
case in select to make non-blocking reads or writes.

The default case makes the send & receieve operations non-blocking
The 1st select statement tries to send data to ch; since ch is a buffered
channel with a capacity of 1, it can hold one value without blocking.
The 2nd select checks if there's data to receieve from ch without blocking.
*/

func nonBlockingChannelOperations() {
	fmt.Println("\nNon-Blocking Channel Operations")

	ch := make(chan int, 1) // buffered channel to allow non-blocking send

	select {
	case ch <- 42:
		fmt.Println("Sent data to channel")
	default:
		fmt.Println("No data to receieve from channel")
	}
}
