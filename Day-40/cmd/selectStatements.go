package main

import "fmt"

func usingSelectStatements() {
	fmt.Println("\nUsing Select Statements")

	intro := `
Use select with a default case to prevent blocking.	
	`
	fmt.Println(intro)

	ch := make(chan int)

	select {
	case ch <- 1:
		fmt.Println("sent value")
	default:
		fmt.Println("No reciever available")
	}
}
