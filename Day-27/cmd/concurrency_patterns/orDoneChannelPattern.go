package concurrencypatterns

import (
	"fmt"
)

/*
This pattern allows a function to stop waiting for a channel’s data when a done signal is received.
*/

func or(done <-chan struct{}, c <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if !ok {
					return
				}
				out <- v
			}
		}
	}()
	return out
}

func OrDone() {
	fmt.Println("\nConcurrency Patterns: Or-Done Channel Pattern")

	done := make(chan struct{})
	c := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			c <- i
		}
		close(c)
	}()

	for val := range or(done, c) {
		fmt.Println(val)
		if val == 3 {
			close(done)
		}
	}

}
