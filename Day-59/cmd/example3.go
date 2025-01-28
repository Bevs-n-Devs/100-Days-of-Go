package main

import "fmt"

func Example3() {
	var scenario = `
	Example 3: Favour Composition Over Inheritance
		- Go doesn't support inheritance, but encourages composition with structs and interfaces.
	
	Code Example:
	type Engine struct {
		Horsepower int
	}

	type Car struct {
		Brand  string
		Engine Engine // Embedding Engine struct
	}
	`
	fmt.Println(scenario)
}
