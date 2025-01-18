package main

import "fmt"

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nYAGNI (You Aren't Gonna Need It) Principles in SWD & Go")

	var definition = `
		Definition:
YAGNI is a principle of software development that emphasizes avoiding unnecessary features or functionality in your code. 
It means you should only implement what is required for your application right now, not what you think might be needed in the future.
	`
	fmt.Println(definition)

	var whyYAGNI = `
		Why YAGNI is Important?
	1. Reduces Complexity: Unnecessary features can complicate the codebase and make it harder to maintain.
	2. Saves Tome: Building features "just in case" wastes development time that could be spent on solving actual problems.
	3. Improces Focus: Encourages focusing on immediate requirements and delivering value to users.
	4. Avoids Over-Engineering: Prevents over-engineering solutions to problems that may never arise.
	5. Avoids Wasted Effort: Requirements often change, making speculative features obsolete.
	`
	fmt.Println(whyYAGNI)

	var keyIdea = `
		Key Idea: Solve Today's Poblem
	When implementing a feature always ask:
		- Does the feature solve a current problem?
		- Is there a confirmed need for it in the near furture?
	If the amder is "no", then don't build it yet.
	`
	fmt.Println(keyIdea)

	Examople1()

	Example2()

	Example3()

	var misappliedYAGNI = `
	When YAGNI Can Be Misapplied:
		- Ignoring Necessary Planning: YAGNI doesn't mean skipping essential design and planning. 
		  It's about not implementing features prematurely, not neglecting necessary preparation.
		- Not Adding Obvious Essentials: If something is clearly required soon (e.g user authentication for an app storing sensitive data),
		  it's not YAGNI to implement it now.
	`
	fmt.Println(misappliedYAGNI)

	var praticalTips = `
	Practical Tips for Applying YAGNI:
		1. Focus on Requirements: Always start with what is explicitly needed for the user or project.
		2. Interative Development: Build features incrementally based on feedback and real needs.
		3. Avoid Speculative Cdoe: Don't write code for hypothetical use cases.
		4. Refactor When Needed: Write simple code now and refactor later as requirements evolve.
		5. Ask "Why Now?": If you are considering adding a current feature, ask whether it's solving a current problem.
		6. Focus on MVP: Aim to deliver a Minimum Viable Product (MVP) with essential features first.
	`
	fmt.Println(praticalTips)

	var finalThought = `
	Final Thought:
	YAGNI is about striking a balance: don't sacrifice simplicity and time for hypothetical future needs.
	Focus on delivering immediate value, and let actual user feedback guide the evolution of your codebase/software.
	`
	fmt.Println(finalThought)
}
