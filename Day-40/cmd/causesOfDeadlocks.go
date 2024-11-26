package main

import "fmt"

func commonCausesOfDeadlocks() {
	fmt.Println("\nCommon Causes of Deadlocks")

	intro := `
We will cover the 3 main causes of deadlocks:
1. Unbuffered Channel Blocking
2. Buffered Channel Mismanagement
3. Improper Use of Mutexes
4. Circular Dependencies
	`
	fmt.Println(intro)

	unbufferedChannels()

	bufferedChannelMismanagement()

	improperUseofMutexes()

	// uncomment to see deadlock
	circularDependencies()
}
