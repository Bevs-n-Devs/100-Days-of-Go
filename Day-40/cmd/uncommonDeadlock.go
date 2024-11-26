package main

import "fmt"

func uncommonCasesOfDeadlocks() {
	fmt.Println("\nUncommon Cases of Deadlocks")

	intro := `
We will cover the 2 uncommon causes of deadlocks:
1. Miconfigured Worker Pools
2. select Statements with Blocking Cases	
	`
	fmt.Println(intro)

	misconfiguredWorkerPools()

	// uncomment to see deadlock
	// selectStatementsWithBlockingCases()
}
