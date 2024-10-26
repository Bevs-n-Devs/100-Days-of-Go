package main

import (
	"fmt"
	"time"
)

/*
Golang select statement is like the switch statement, which is used for multiple channels operation.
This statement blocks until any of the cases provided are ready.

The select statement is powerful for coordinating Goroutines and working with multiple channels,
especially when we want to implement timeouts, non-blocking operations, or respond to multiple channels.

The select statement in Go allows a Goroutine to wait on multiple channel operations simultaneously.
You can think of it as a switch statement but for channels. It selects a case that’s ready to
proceed and executes it, which is very helpful when working with multiple channels

select {
case <-channel1:
     do something when channel1 has data
case msg := <-channel2:
     do something with msg when channel2 has data
default:
     execute if none of the channels are ready (non-blocking case)
}

Each case is a channel operation (send or receive). The select statement will block until one of the cases can proceed.

The select statement itself does not automatically defer or block other goroutines; it only blocks until one of the cases in the select is ready to proceed.
So, it doesn't provide synchronization for other resources (like shared data) on its own.

The sendMessage function sends a message to the channel after a delay.
select waits on both channels, and whichever channel receives data first is selected.
Since channel2 has a shorter delay, its message is printed first.
*/

func sendMessage(channel chan string, message string, delay time.Duration) {
	time.Sleep(delay)
	channel <- message
}

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nUsing Select for Basic Channel Operations")

	channel1 := make(chan string)
	channel2 := make(chan string)

	go sendMessage(channel1, "Message from Channel 1", 2*time.Second)
	go sendMessage(channel2, "Message from Channel 2", 1*time.Second)

	select {
	case msg1 := <-channel1:
		fmt.Println(msg1)
	case msg2 := <-channel2:
		fmt.Println(msg2)
	}

	selectTimeout()

	nonBlockingChannelOperations()

	concurrentTasksWithSelect()

	timeoutRetryMechanism()
}

/*
SELECT SUMMARY:

Select Statement: Allows a Goroutine to wait on multiple channel operations.
				  It picks the first ready case to proceed, which is ideal for managing multiple concurrent tasks.

Timeouts with time.After: The time.After function creates a channel that sends a value after
						  a specified duration, making it easy to implement timeouts.

Non-Blocking Operations: By using select with a default case, we can perform non-blocking channel
						 operations to check channel states without waiting.

Real-World Usage: Using select with Goroutines allows you to handle concurrent
				  operations with retries, task coordination, and more.
*/
