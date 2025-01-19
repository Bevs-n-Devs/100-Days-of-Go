package main

import "fmt"

func main() {
	fmt.Println("Hello world, hello Yaw!")
	fmt.Println("\nKISS (Keep It Simple, Stupid) Principles in SWD & Go")

	var definition = `
The KISS principle advocates for simplicity in code, architecture and design. 
The idea is to keep solutions straightforward and avoid unnecessary complexity.
This makes the system easier to understand, maintain and extend.
	`
	fmt.Println(definition)

	var whyKISS = `
	Why KISS Matters
	1. Improves Readability: Simple code is easier to read and understand for developers.
	2. Facilitates Maintenance: Debugging and modifying simple code is less error-prone.
	3. Speeds Up Development: Developers spend less time trying to understand or navigate complex code.
	4. Reduces Bugs: Complexity breeds errors, so simpler code is less likely to contain bugs.
	`
	fmt.Println(whyKISS)

	var keyIdeas = `
	Key Ideas
	1. Solve problems using the simplest solution that works.
	2. Avoid overengineering or overgeneralizing.
	3. Don't add unnecessary abstractions or layers unless they provide a clear value.
	`
	fmt.Println(keyIdeas)

	Example1()

	Example2()

	Example3()

	Example4()

	var tips = `
	Tips for Applying KISS
	1. Start with the simplest solution: Only add complexity if required.
	2. Use Go idioms: Follow Go's philosophy of simplicity and readability.
	3. Avoid unnecessary abstractions: Don’t add interfaces or layers unless they’re needed.
	4. Refactor over time: Simplicity doesn't mean sloppy. Refactor as requirements evolve.
	5. Think about future maintainers: Write code as if the next person reading it has no context.

	By following the KISS principle, you'll write code that is easier to maintain, extend, and debug while avoiding unnecessary complexity.
	`
	fmt.Println(tips)
}
