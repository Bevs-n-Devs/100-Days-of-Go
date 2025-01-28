package main

import "fmt"

func Example7() {
	var scenario = `
	Example 7: Write Small, Focused Functions
		- Functions in Go ideally do one thing well and return quickly
	
	Code Example:
	// Small, focused function
	func IsEven(n int) bool {
		return n%2 == 0
	}
	`
	fmt.Println(scenario)
}
