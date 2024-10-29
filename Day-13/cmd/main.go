package main

import "fmt"

/*
Today, we’ll dive into interface-based testing. Testing code that interacts with external resources (like databases or APIs)
can be challenging. Using interfaces to abstract dependencies makes it possible to test code without relying on the actual
resources. You’ll learn how to create interfaces, mock dependencies, and write unit tests that use these mocks.
*/

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nInterface-Based Testing")
}

/*
BEST PRACTICES

1. Keep Interfaces Small: Avoid large interfaces. Smaller, single-method interfaces are more flexible.
2. Mock External Dependencies: Use mocks for resources like databases, APIs, or third-party services.
3. Verify Mock Calls When Needed: Use libraries like Testify to verify that mocks are used correctly in your tests.
4. Document Expected Behavior: Write comments in test cases that clarify each expected outcome.
*/
