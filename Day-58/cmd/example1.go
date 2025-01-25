package main

import "fmt"

func Example1() {
	fmt.Println("\nExample 1: Avoiding Long Chains")
	var scenario = `
	Violation of PoLA:
	order.User.Address.City

	Following PoLA:
		- Add intermediary methods to prevent directly accessing deeply nested structures

	Code Example:
	func (u *User) GetCity() string {
		return u.Address.City
	}

	func (o *Order) GetCity() string {
		return o.User.GetCity()
	}
		- Now the client only needs to interact with Order
	`
	fmt.Println(scenario)
}
