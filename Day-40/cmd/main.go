package main

import "fmt"

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nDeep Dive into Deadlocks and Deadlock Prevention in Go")

	intro := `
Deadlocks are one of the most common and subtle problems in concurrent programming. 
They occur when two or more goroutines are waiting for each other to release resources, and none of them can proceed.

In this lesson, we will thoroughly explore:
	1. What is a Deadlock?
	2. Common Causes of Deadlocks
	3. Uncommon Causes of Deadlocks
	4. How to Detect Deadlocks
	5. Prevention Strategies
	6. Advanced Considerations
	`
	fmt.Println(intro)

	// uncomment to see the runtime panic
	// whatIsDeadlock()

	commonCausesOfDeadlocks()

	uncommonCasesOfDeadlocks()

	preventionStrategies()

	advancedConsiderations()

	conclusion := `
	Conclusion:
	
Deadlocks are a critical issue in concurrent programming, and understanding their causes and prevention is essential for 
building robust Go applications. By following best practices—avoiding circular dependencies, using buffered channels wisely, 
and employing tools like select, sync.WaitGroup, and context.Context—you can write efficient and deadlock-free code.

With this knowledge, you are equipped to tackle complex concurrency challenges in Go.
	`
	fmt.Println(conclusion)
}
