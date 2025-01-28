package main

import "fmt"

func Example8() {
	var scenario = `
	Example 8: Follow Go Formatting and Conventions
	Use the Go tools (gofmt, golangci-lint) to ensure your code is properly formatted and free of lint issues. The Go community follows specific conventions:
		- No semicolons (;) at the end of lines.
		- Use tabs (not spaces) for indentation.
	
	Code Example:
	// Idiomatic Go formatting
	func Hello(name string) string {
		return fmt.Sprintf("Hello, %s!", name)
	}
	`
	fmt.Println(scenario)
}
