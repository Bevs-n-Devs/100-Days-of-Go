package main

import "fmt"

func Example1() {
	var scenario = `
	Example 1: Simplicity is Key
		- Write clear, simple, and concise code.
		- Avoid overengineering and unnecessary abstraction.

	Overengineering Code Example:
	func CalculateSum(numbers []int) (sum int, err error) {
		if numbers == nil {
			return 0, fmt.Errorf("numbers slice is nil")
		}
		for _, n := range numbers {
			sum += n
		}
		return
	}

	Idiomatic Go Example:
	func Sum(numbers []int) int {
		sum := 0
		for _, n := range numbers {
			sum += n
		}
		return sum
	}
	`
	fmt.Println(scenario)
}
