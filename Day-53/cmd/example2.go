package main

import "fmt"

func Example2() {
	var scenario = `
Example 2: Prematurely Generalizing Code

	Scenario:
	You're writing a function to calculate the total price of items in a shopping cart.

	Premature Implementation:
	You write a generic, overengineered function that supports:
		- Multiple currencies
		- Discounts
		- Tax rates for multiple countries

	Code Example:
	type Item struct {
		Name     string
		Price    float64
		Quantity int
		Currency string
		Discout  float64
		TaxRate  float64
	}
	
	func CalculateTotal(items []Item) float64 {
		total := 0.0
		for _, item := range items {
			total += (item.Price * item.Quantity) * (1 - item.Discount) * (1 + item.TaxRate)
		}
		return total
	}

	Problems:
		- Overengineering: You don't even support multiple currencies ot tax rates yet.
		- Complexity: Testing becomes harder with all these additional features.
		- Missed Deadlines: Adding inneccessary functionality delays the delivery of the actual feature.

	YAGNI Approach:
	Implement only the functionality you need right now. If all items are in one currency and taxes/discounts are not
	yet needed, simplify the function to calculate the total price without these features.

	Code Example:
	func CalculateTotal(items []Item) float64 {
		total := 0.0
		for _, item := range items {
			total += item.Price * float64(item.Quantity)
		}
		return total
	}
	
	If the requirements change, you can refactor and extend the function later.
	`
	fmt.Println(scenario)
}
