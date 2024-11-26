package main

import (
	"context"
	"fmt"
	"time"
)

func cancellationWithContext() {
	fmt.Println("\nCancellation with Context")

	intro := `
Use context.Context to cancel long-running or unresponsive operations.	
	`
	fmt.Println(intro)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ch := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- 42
	}()

	select {
	case val := <-ch:
		fmt.Println("Recievedv", val)
	case <-ctx.Done():
		fmt.Println("Operation timed out")
	}
}
