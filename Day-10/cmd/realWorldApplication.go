package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
We will create a small program that retries a task if it fails to complete within a certain timeout.
This can be useful in network calls where you might want to retry connecting to a service if the response doesn’t arrive in time.

unreliableTask simulates a task with a random chance of success or failure.
In main, we use select to wait for the task’s result or timeout after 3 seconds.
If the task times out, the loop retries up to 3 times.
*/

func unreliableTask(ch chan string) {
	// simulate random success or failure
	if rand.Intn(2) == 0 {
		time.Sleep(2 * time.Second) // simulate delay on success
		ch <- "Task succeeded"
	} else {
		time.Sleep(4 * time.Second) // simulate longer delay on failure
	}
}

func timeoutRetryMechanism() {
	fmt.Println("\nReal-World Application: A Simple Timeout Retry Mechanism")

	rand.Seed(time.Now().UnixNano()) // seed random number generator

	retries := 3
	success := false

	for i := 0; i < retries; i++ {
		ch := make(chan string)

		go unreliableTask(ch)

		select {
		case result := <-ch:
			fmt.Println(result)
			success = true
			break
		case <-time.After(3 * time.Second):
			fmt.Println("Attempt", i+1, "timed out. Retrying...")
		}

		if success {
			break
		}
	}

	if !success {
		fmt.Println("Task failed after all retries.")
	}

}
