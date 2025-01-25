package main

import "fmt"

func Example2() {
	fmt.Println("\nExample 2: Avoiding Exposing Internal State")
	var scenario = `
	Violation of PoLA:
		- Returning a mutable internal state breaks encapsulation

	Code Example:
	func (u *User) GetAddress() *Address {
		return u.Address // Violates LoD by exposing internal details
	}

	Following PoLA:
		- Provide methods that expose only what’s needed
	
	Code Example:
	func (u *User) GetCity() string {
		return u.Address.City
	}
	`
	fmt.Println(scenario)
}
