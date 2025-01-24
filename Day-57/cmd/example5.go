package main

import "fmt"

func Example5() {
	fmt.Println("\nExample 5: Intuitive Struct Design")
	var scenario = `
	Violation of PoLA:
		- This Counter struct allows unexpected negative counts.

	Code Example:
	type Counter struct {
		count int
	}

	func (c *Counter) Decrement() {
		c.count-- // Unexpected: Might allow negative counts
	}

	func (c *Counter) Value() int {
		return c.count
	}

	Following PoLA:
	func (c *Counter) Decrement() error {
		if c.count == 0 {
			return fmt.Errorf("counter cannot be negative")
		}
		c.count--
		return nil
	}
	`
	fmt.Println(scenario)
}
