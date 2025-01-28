package main

import "fmt"

func Example6() {
	var scenario = `
	Example 6: Handle Errors Where They Occur
		- Don't ignore errors, and handle them as close to their origin as possible.

	Avoid Ignoring Errors:
	// BAD: Ignoring errors
	data, _ := os.ReadFile("config.json") // Ignoring the error is not idiomatic

	Correct Approach:
	// GOOD: Handle the error
	data, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatalf("could not read config: %v", err)
	}
	`
	fmt.Println(scenario)
}
