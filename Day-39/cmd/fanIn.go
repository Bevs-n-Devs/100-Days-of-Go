package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- string, msg string, delay time.Duration) {
	for i := 0; i < 3; i++ {
		ch <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(delay)
	}
	close(ch)
}

func fanIn() {
	fmt.Println("\nFan In")

	intro := `
Fan-In: Multiple channels write to the same channel
	`
	fmt.Println(intro)

	ch1 := make(chan string)
	ch2 := make(chan string)

	go producer(ch1, "Hello", 1*time.Second)
	go producer(ch2, "World", 2*time.Second)

	// while loop
	for {
		select {
		case msg1, ok := <-ch1:
			if ok {
				fmt.Println(msg1)
			} else {
				ch1 = nil
			}
		case msg2, ok := <-ch2:
			if ok {
				fmt.Println(msg2)
			} else {
				ch2 = nil
			}
		}

		// stop the loop when both channels are empty or closed
		if ch1 == nil && ch2 == nil {
			break
		}
	}
}
