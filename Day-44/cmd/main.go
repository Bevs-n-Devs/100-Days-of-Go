package main

import "fmt"

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\n	Mastering Dependency Injection in Go")

	intro := `
Dependency Injection (DI) is a critical concept for mid-level developers to understand as it is key to building scalable, testable, and maintainable Go applications. 
It involves providing dependencies to a component rather than the component creating them itself. This pattern helps decouple components and makes testing and scaling much easier.
	`
	fmt.Println(intro)

	whatIsDI := `
	What is Dependency Injection (DI)?

- Dependency Injection is a design pattern where objects receive their dependencies from an external source rather than creating them internally.
- In Go, this is often done via constructor injection, interface injection, or struct embedding.
	`
	fmt.Println(whatIsDI)

	whyUseDI := `
	Why Use Dependency Injection (DI)?

1. Improves Testability: Dependencies can be mocked or replaced during tests.
2. Reduces Coupling: Components are loosely coupled, making it easier to refactor or swap implementations.
3. Enhances Readability: Dependencies are clearly visible in the component's constructor or setup function.
4. Facilitates Scalability: Makes your code more modular and easier to extend.
	`
	fmt.Println(whyUseDI)

	exampleWithoutDI()

	exampleWithDI()

	exampleWithConstructorInjection()

	// uncomment to see dependency injection frameworks example
	// dependencyInjectionFrameworks()

	bestPractices := `
	Best Practices for DI

1. Inject Interfaces, Not Concrete Types:
	- Depend on abstractions (interface) rather than implementations.

	type Service interface {
		DoSomething()
	}

2. Use Constructor Injection:
	- Explicitly inject dependencies using constructors.
3. Avoid Global State:
	- Use dependency injection to manage state instead of global variables.
4. Mock Dependencies for Testing:
	- Provide mock implementations for easier unit testing.
5. Keep Dependencies Minimal:
	- Avoid injecting unnecessary dependencies into components.
	`
	fmt.Println(bestPractices)

	commonPitfalls := `
	Common Pitfalls of DI

1. Over-Injection:
	- Avoid injecting too many dependencies into a single component. This is a sign of poor design.
2. Leaky Abstractions:
	- Ensure that injected interfaces are well-defined and don’t expose unnecessary details.
3. Circular Dependencies:
	- Avoid circular dependencies, which can cause runtime errors or make the code difficult to understand.
	`
	fmt.Println(commonPitfalls)

	whenToUseDI := `
	When to Use Dependency Injection (DI)?

1. When components have complex dependency trees.
2. When mocking is required for unit testing.
3. When building modular, scalable systems (e.g., microservices).
	`
	fmt.Println(whenToUseDI)

	conclusion := `
	CONCLUSION 

Dependency Injection is a fundamental pattern for writing maintainable and scalable Go applications. 
While Go's simplicity means DI frameworks are often unnecessary, understanding DI principles and techniques 
will significantly improve your code quality, making it easier to test, extend, and refactor.
	`
	fmt.Println(conclusion)
}
