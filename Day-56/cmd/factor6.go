package main

import "fmt"

func Factor6() {
	var factor6 = `
	6. Processes: Execute the App as One or More Stateless Processes
		- Principle: App processes should be satteless, with state stored in external systems like database or caches.
		- Why: Enables horizontal scaling and resilience.

		Example:
			- Avoid storing user data in memory. Use redis or another external store instead.
	`
	fmt.Println(factor6)
}
