package main

import "fmt"

func selectStatementsWithBlockingCases() {
	fmt.Println("\nSelect Statements with Blocking Cases")

	intro := `
Improperly handling a select statement without a default case can lead to deadlocks if all channels block.
	`
	fmt.Println(intro)

	ch1 := make(chan int)
	ch2 := make(chan int)

	select {
	case ch1 <- 1: // blocks forever because theres no receiver
	case ch2 <- 2: // blocks forever because theres no receiver
	}
}
