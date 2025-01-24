package main

import "fmt"

func Example3() {
	fmt.Println("\nExample 3: Unexpected State Changes")
	var scenario = `
	Violation of PoLA:
		- A function modifies global state unexpectedly.
	
	Code Example:
	var globalCounter int

	func IncrementAndGetGlobalCounter() int {
		globalCounter++
		return globalCounter
	}

	Problem:
		- The function is performing two actions (incrementing and returning) while changing an external variable, which might surprise users expecting a simple getter.
		
	Following PoLA:
		- Use clear separation of actions and avoid modifying external state unexpectedly

	Code Example:
	func IncrementCounter(counter int) int {
		return counter + 1
	}
	`
	fmt.Println(scenario)
}
