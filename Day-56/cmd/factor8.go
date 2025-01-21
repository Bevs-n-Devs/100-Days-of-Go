package main

import "fmt"

func Factor8() {
	var factor8 = `
	8. Concurrency: Scale Out via the Process Model
		- Principle: Scale horizontally by running multiple instances of lightweight processes.
		- Why: Leverages distributed systems for scaling.

		Example:
			- Use Kubernetes or Docker Compose to manage multiple instances.

		Code Example:
			replicas: 3
	`
	fmt.Println(factor8)
}
