package main

import "fmt"

func Example5() {
	var scenario = `
	Example 5: Keep Error Handling Simple
		- Error handling is explicit in Go. Use the error type and avoid nested error checks where possible.
	
	Code Example:
	func ReadConfig(filename string) ([]byte, error) {
		data, err := os.ReadFile(filename)
		if err != nil {
			return nil, fmt.Errorf("failed to read config: %w", err)
		}
		return data, nil
	}
	`
	fmt.Println(scenario)
}
