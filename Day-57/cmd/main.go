package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println("Hello world, hello Yaw!")
	fmt.Println("\nThe Principle of Least Ashtonishmet - PoLA")

	var definition = `
The Principle of Least Ashtonishmet (PoLA) states that software interfaces, APIs and code
should behave in ways that are predictable and intuitive to users, developers, or ther stakeholders.
The goal is to minimize surprises and align with other user or developer expectations.	
	`
	fmt.Println(definition)

	var whyPoLA = `
	Why PoLA Matters
		1. Improves Usability: Predictable behavior makes software easier to use.
		2. Reduces Errors: Users and developers are less likely to make mistakes if behavior matches expectations.
		3. Enhances Maintainability: Developers can more easily understand and modify code that behaves predictably.
		4. Builds Trust: Intuitive interfaces encourage users and developers to feel confident in the software.
	`
	fmt.Println(whyPoLA)

	var coreConcepts = `
	Core Concepts of PoLA
		1. Consistency: Follow established conventions and patterns.
		2. Simplicity: Avoid overcomplicating interactions or behaviors.
		3. Transparency: Make the behavior of the system clear and easy to reason about.
		4. Align with Expectations: Match the mental model of your users or developers.
	`
	fmt.Println(coreConcepts)

	Example1()
	Example2()
	Example3()
	Example4()

	var guidlines = `
	Guidelines for Applying PoLA
		1. Follow Established Conventions:
			- Adhere to language or framework idioms.
			- Respect common patterns in user interfaces (e.g., placing "OK" on the right and "Cancel" on the left).
		2. Make Behavior Predictable:
			- Avoid side effects in functions unless explicitly stated.
			- Ensure APIs do what their names imply.
		3. Use Clear and Descriptive Names:
			- Variables, functions, and classes should have names that describe their purpose accurately.
		4. Document Exceptions:
			- If behavior deviates from the norm, document it clearly.
		5. Test for Intuitiveness:
			- Gather feedback from users and developers to ensure interfaces match expectations.
	`
	fmt.Println(guidlines)

	Example5()
	Example6()

	var benefits = `
	Benefits of Following PoLA
		1. Reduces Cognitive Load: Users and developers don’t have to guess how things work.
		2. Prevents Errors: Predictable behavior reduces the chance of misusing code or interfaces.
		3. Encourages Adoption: Intuitive tools and systems are easier to learn and adopt.
	`
	fmt.Println(benefits)

	var conclusion = `
By following the Principle of Least Astonishment, you ensure that your software and APIs are intuitive, 
reliable, and enjoyable to use, fostering trust and ease for both users and developers.
	`
	fmt.Println(conclusion)
}
