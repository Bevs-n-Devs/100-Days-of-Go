package main

import "fmt"

func Example10() {
	var scenario = `
	Example 10: Use Go's Concurrency Features
		- Go has first-class support for concurrency via goroutines and channels.
	
	Code Example (Goroutines & Channels):
	func Process(data []int, results chan<- int) {
		for _, v := range data {
			results <- v * v
		}
		close(results)
	}

	func main() {
		data := []int{1, 2, 3, 4, 5}
		results := make(chan int)

		go Process(data, results)

		for res := range results {
			fmt.Println(res)
		}
	}
	`
	fmt.Println(scenario)
}
