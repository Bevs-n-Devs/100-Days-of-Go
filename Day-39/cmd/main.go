package main

import "fmt"

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nDeep Dive Into Channels")

	intro := `
Channels are one of the most powerful concurrency primitives in Go. 
They enable communication between goroutines, making it easier to share data and coordinate tasks in a concurrent environment.
	`
	fmt.Println(intro)

	whatAreChannels := `
	What are Channels in Go?

A channel in Go is a mechanism through which goroutines can communicate by sending and receiving values of a specified type. 
Channels are strongly typed, meaning they can only transfer values of a single type.
	`
	fmt.Println(whatAreChannels)

	declareAndCreate := `
	Declaring and Creating Channels

To create a channel, you can use the make() function. 
The make() function takes two arguments: the type of the channel and the capacity of the channel.

	ch := make(chan int) // creates a channel for integers
	`
	fmt.Println(declareAndCreate)

	sendRecieve := `
	Basic Syntax for Sending and Receiving Values

- Send: ch <- value (Send value into the channel ch.)
- Receive: value := <-ch (Receive a value from the channel ch.)	
	`
	fmt.Println(sendRecieve)

	howChannelsWork := `
	How Channels Work

Channels act as conduits between goroutines:
1. A sending goroutine places a value into a channel.
2. A receiving goroutine retrieves the value from the channel.
3. Synchronization: Channels block the sending and receiving goroutines until the operation is complete.
	`
	fmt.Println(howChannelsWork)

	unbufferedChannels()

	bufferedChannels()

	closingChannels()

	channelDirections()

	selectStatement()

	channelUseCases()

	bestPractices := `
	BEST PRACTICES

1. Avoid Deadlocks
	- Always ensure that every send operation has a corresponding receive.
	- Be cautious when using unbuffered channels, as they block both the sender and receiver.
2. Close Channels Gracefully
	- Use close() to signal that no more data will be sent.
	- Avoid closing a channel from multiple goroutines.
3. Prefer Buffered Channels for Non-Blocking Operations
	- Buffered channels provide flexibility when the producer and consumer run at different speeds.
4. Use Select for Multiple Channel Operations
	- select simplifies working with multiple channels.
5. Leverage Directional Channels
	- Restricting channel directions improves code clarity and prevents misuse.
	`
	fmt.Println(bestPractices)

	conclusion := `
	Conclusion

Channels are a core feature of Go's concurrency model, enabling goroutines to communicate safely and effectively. 
By mastering their use—unbuffered vs. buffered, closing channels, select statements, and advanced patterns like 
worker pools and fan-out/fan-in—you can build powerful, concurrent applications in Go.
	`
	fmt.Println(conclusion)
}
