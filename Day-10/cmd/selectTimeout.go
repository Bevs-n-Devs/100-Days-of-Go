package main

import (
	"fmt"
	"time"
)

/*
In real-world applications, you may need to set a timeout for a channel operation if it takes too long to receive a response.

The time.Afetr function creates a channel that sends the curent time afgter the specified duration.
If the ch channel doesn't receieve data within 2 secs, the timeout case executes, preventing
the program from waiting indefinately.
*/

func selectTimeout() {
	fmt.Println("\nUsing Select with Timeouts")

	ch := make(chan string)
	// create anonyous Goroutine
	go func() {
		time.Sleep(time.Second * 3)
		ch <- "data from channel"
	}()

	select {
	case msg := <-ch:
		fmt.Println("Receieved:", msg)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout reached before receiving data")
	}
}
