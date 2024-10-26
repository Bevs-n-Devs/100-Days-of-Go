package main

import (
	"fmt"
	"time"
)

/*
Using select, you can manage multiple tasks that run in parallel.
This is often used in applications like web servers to handle multiple client requests simultaneously.

Two tasks are running concurrently with different durations.
Each task sends a message to the ch channel upon completion.
The for loop with select reads from the channel twice to receive both messages as they complete.
*/

func task(id int, duration time.Duration, ch chan<- string) {
	time.Sleep(duration)
	ch <- fmt.Sprintf("Task %d completed", id)
}

func concurrentTasksWithSelect() {
	fmt.Println("\nConcurrent Tasks With Select")

	ch := make(chan string)

	go task(1, 2*time.Second, ch)
	go task(2, 1*time.Second, ch)

	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch:
			fmt.Println(msg)
		}
	}

}
