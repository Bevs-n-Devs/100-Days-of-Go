package main

import "fmt"

func Factor11() {
	var factor11 = `
	11. Logs: Treat Logs as Event Streams.
		- Principle: Apps should output logs as streams to stdout and stderr, leaving log processing to the environment.
		- Why: Enables centralized log aggregation and analysis.
		
		Code Example:
			log.Println("App started successfully!")
	`
	fmt.Println(factor11)
}
