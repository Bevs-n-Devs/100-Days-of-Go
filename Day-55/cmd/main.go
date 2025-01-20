package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println("Hello world, hello Yaw!")
	fmt.Println("\nSeparation of Concerns")

	var definition = `
Separation of Concern is a design principle that ionvolves dividing a program into distinct,
independent saections, each responsible for a single functionality or concern. This reduces 
coupling (dependencies between parts) and improves modularity, maintainability and scalability.
	`
	fmt.Println(definition)

	var whySOC = `
	Why SOC Matters
		1. Improve Code Readability: Each module or section has a clear responsibility, making easier to understand.
		2. Eases Maintenance: Changes to one concern don't unintentionally impact others.
		3. Encourages Reusability: Independent components can be reused across different parts of the program.
		4. Simplifies Testing: Testing isolated concerns is simpler than testing tightly coupled code.
	`
	fmt.Println(whySOC)

	var coreIdea = `
	Core Idea
		- Each corncern has its own responsibility.
		- Concerns should interact through well-defined boundaries.
		- Avoid "doing too much" in a single place.
	`
	fmt.Println(coreIdea)

	Example1()

	Example2()

	var tips = `
	Tips for Applying SAeparating of Concerns
		1. Understanding Responsibilities: Identify distinct responsibilities in your program (eg. database access, buisiness logic, user interface).
		2. Use Layers: Divide your code into logical layers (eg. controller, service, repository).
		3. Define Interfaces: Use interfaces to abstract interactions between layers (eg. repository interfaces for database access).
		4. Avoid Coupling: Minimize direct dependencies between unrelated concerns.
		5. Focus on Testability: Write code that allows each layer or concern to be tested independently.
	`
	fmt.Println(tips)

	var realWorldExamples = `
	Real-World Applications
		1. Web Applications: Use MVC (Model-View-Controller) pattern or similar patterns.
		2. Distributed Systems: Separate services for by domain - authentication, payments, notifications etc).
		3. Microservices: Each service handles a specific concern, communicating via APIs.
		
By folloing the Separation of Concerns, you can build systems that are easier to understand, modify and scale, reducing the risk of cascading changes and bugs.
	`
	fmt.Println(realWorldExamples)

}
