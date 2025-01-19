package main

import "fmt"

func Example1() {
	fmt.Println("\nExample 1: Over-Complicating Something Simple")

	var scenario = `
	Scenario:
	You need a function to reverse a string.

	Over-Complicated Approach:
	Using an abstract and over-engineered solution

	Code Example:
	func ReverseString(input string) (output string, err error) {
		if input == "" {
			return "", fmt.Errorf("input string cannot be empty")
		}

		runes := []rune(input)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		output = string(runes)
		return output, nil
	}
	
	Problems:
		1. Unnecessary error handling for reversing a string.
		2. Add complexity with no clear benefit.

	KISS Approach:
	Here's a simpler solution that works:

	Code Example:
	func ReverseString(input string) string {
		runes := []rune(input)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes)
	}

	Benefits:
		- No unnecessary error handling.
		- Straightforward and easy to understand.
	`
	fmt.Println(scenario)
}
