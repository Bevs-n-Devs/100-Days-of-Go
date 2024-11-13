package main

import (
	"fmt"
)

/*
Topics to cover:

1. Basics of Goroutines
2. Goroutines with Channels
3. Concurrency Patterns
4. Context for Cancellation and Timeouts
5. Avoiding Common Pitfalls
6. Best Practices and Optimization


*/

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nDeep Dive Into Gorotines")
	fmt.Println(`
In this lesson we will cover the following:

1. Basics of Goroutines
2. Goroutines with Channels
3. Concurrency Patterns
4. Context for Cancellation and Timeouts
	`)

	goroutineBasics()

	goroutineChannel()

	concurrencyPatterns()

}

/*
BEST PRACTICES
1. Use Goroutines Sparingly: Only use goroutines when parallel execution provides an actual performance benefit.
2. Limit Channel Usage: Channels are powerful, but they can introduce complexity. Use direct function calls or mutexes when simpler.
3. Avoid Overuse of go Keyword: Each go introduces a new concurrent function call. Excessive usage can degrade performance.
*/
