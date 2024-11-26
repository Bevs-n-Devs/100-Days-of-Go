package main

import "fmt"

func gracefulChannelClosing() {
	fmt.Println("\nGraceful Channel Closing")

	intro := `
Always close channels when no more values will be sent. Use a for range loop for safe reading.
	`
	fmt.Println(intro)

	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	for val := range ch {
		fmt.Println(val)
	}
}
