package main

import "fmt"

func advancedConsiderations() {
	fmt.Println("\nAdvanced Considerations")

	intro := `
Advanced Considerations:
1. Timeouts & Contexts
2. Cancellation with Context
	`
	fmt.Println(intro)

	timeoutsAndContexts()

	cancellationWithContext()
}
