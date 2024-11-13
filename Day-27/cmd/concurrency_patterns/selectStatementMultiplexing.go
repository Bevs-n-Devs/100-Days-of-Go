package concurrencypatterns

import (
	"fmt"
	"time"
)

/*
The select statement listens on multiple channels and performs actions based on whichever channel is ready first.

select statementrs are like swicth statements for channels!
*/

func SelectStatements() {
	fmt.Println("\nConcurrency Patterns: Select Statements For Multiplexing")

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

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout")
	}

}
