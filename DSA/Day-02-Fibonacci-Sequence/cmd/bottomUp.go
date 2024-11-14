package main

func fibonacciIterative(n int) int {
	if n <= 1 {
		return n
	}

	// assign first two positions
	prev1, prev2 := 0, 1

	// start the loop from 2nd position
	for i := 2; i <= n; i++ {
		current := prev1 + prev2 // current position
		prev1 = prev2            // prev1 moves up by one
		prev2 = current          // prev2 updated to current position
	}

	return prev2 // return the answer
}
