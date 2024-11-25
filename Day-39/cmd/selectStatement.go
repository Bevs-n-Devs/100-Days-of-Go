package main

import (
	"fmt"
	"time"
)

func selectStatement() {
	fmt.Println("\nSelect Statement")

	intro := `
Select Statement
The select statement is used to wait on multiple channel operations, enabling more flexible communication patterns.	
	`
	fmt.Println(intro)

	// create channels
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Message from ch1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Message from ch2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
	//
	expanation := `
Exponation: 
- The select statement listens for data from multiple channels.
- Whichever channel is ready first will be processed.
	`
	fmt.Println(expanation)
}
