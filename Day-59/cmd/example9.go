package main

import "fmt"

func Example9() {
	var scenario = `
	Example 9: Use Interfaces to Define Behaviour
		- Interfaces in Go are designed to capture behaviour, not data structure.

	Code Example:
	// Idiomatic use of interface
	type Shape interface {
		Area() float64
	}

	type Circle struct {
		Radius float64
	}

	func (c Circle) Area() float64 {
		return math.Pi * c.Radius * c.Radius
	}
	`
	fmt.Println(scenario)
}
