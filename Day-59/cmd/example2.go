package main

import "fmt"

func Example2() {
	var scenario = `
	Example 2: Use Go's Built-in Features
		- Leverage Go’s built-in features like slices, maps, channels, goroutines, and defer. Avoid reinventing the wheel.

	Idomatic Code Example:
	func ReadFile(filename string) ([]byte, error) {
		file, err := os.Open(filename)
		if err != nil {
			return nil, err
		}
		defer file.Close() // Ensures file is closed when function exits
		return io.ReadAll(file)
	}
	`
	fmt.Println(scenario)
}
