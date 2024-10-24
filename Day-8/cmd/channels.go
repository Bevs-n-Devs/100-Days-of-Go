package main

import "fmt"

/*
Goroutines run in the same address space, so access to shared memory must be synchronized

Channels are a typed conduit through which you can send and receive values with the channel operator, <-

(eg.)
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and assign value to v

(The data flows in the direction of the arrow.)
Like maps and slices, channels must be created before use:

ch := make(chan int)

By default, sends and receives block until the other side is ready.
This allows goroutines to synchronize without explicit locks or condition variables
*/

func sum(s []int, c chan int) {
	sum := 0

	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to channel
}

func channels() {
	fmt.Println("\nUnderstanding Channels")

	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c // recieve from c

	add := x + y
	fmt.Printf("%d + %d = %v\n", x, y, add)
}
