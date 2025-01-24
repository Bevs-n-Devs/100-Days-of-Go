package main

import "fmt"

func Example1() {
	fmt.Println("\nExample 1: Naming Functions and Methods")
	var scenario = `
	Violation of PoLA:
		- A function name that doesn't match its behaviour.
	
	Code Example:
	// Looks like it calculates a square but actually cubes the number
	func Square(num int) int {
		return num * num * num
	}

	Problem:
		- This surprises developers because the name Square implies squaring a number, not cubing it.


	PoLA Code Example:
	func Cube(num int) int {
		return num * num * num
	}
		- The name clearly reflects what the function does, making it intuitive.
	`
	fmt.Println(scenario)
}
