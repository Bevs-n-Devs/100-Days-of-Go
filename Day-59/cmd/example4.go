package main

import "fmt"

func Example4() {
	var scenario = `
	Example 4: Use Short Variable Names for Small Scopes
		- Use descriptive names for global variables but short names (like i, x, err) for local ones
	
	Code Example:
	for i, v := range numbers {
		fmt.Println(i, v)
	}
	`
	fmt.Println(scenario)
}
