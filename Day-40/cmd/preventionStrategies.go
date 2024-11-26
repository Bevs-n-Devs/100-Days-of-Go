package main

import "fmt"

func preventionStrategies() {
	fmt.Println("\nPrevention Strategies")

	intro := `
Prevention Stategies:
1. Design Principles
2. Graceful Channel Closing
3. Use Select Statements with Default Cases
4. Use sync.WaitGroup
5. Avoid Shared State	
	`
	fmt.Println(intro)

	designPrinciples := `
Design Principles:
1. Avoid Circular Dependencies: Always acquire locks in a consistent order.

	mu1, mu2 := &sync.Mutex{}, &sync.Mutex{}

	lockInOrder := func() {
		mu1.Lock()
		defer mu1.Unlock()

		mu2.Lock()
		defer mu2.Unlock()
	}

2. Use Buffered Channels Judiciously: Buffer size should accommodate expected workloads.
3. Avoid Long-Lived Locks: Keep critical sections (protected by sync.Mutex) as short as possible.
	`
	fmt.Println(designPrinciples)

	gracefulChannelClosing()

	usingSelectStatements()

	syncWaitGroup()

	sharedState()
}
