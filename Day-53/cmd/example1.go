package main

import "fmt"

func Examople1() {
	var scenario = `
Example 1: Prematurely Adding Complexity

	Scenario:
	You're building a to-do list application with the basic requirements:
		1. Add a task
		2. Mark a task as complete
		3. Remove a task
	
	Premature Implementation:
	You decide to add features for categorizing tasks into projects, tagging, assigning priorities, and sharing tasks with other users.
	None of these features are currently required, but you think they might be useful in the future.

	Problems:
		- Waste of Time & Effort: If users don't need these features, you've wasted development time building them.
		- Increased Maintenance: More features mean more code to maintain and test.
		- Delay in Delivering Core Features: Time spent on "extra" features delays delivering what users actually need.

	YAGNI Approach:
	Focus on implementing the basic requirements first (add, mark as complete, delete).
	Only addd the extra features when there's a confirmed need for them.
	`
	fmt.Println(scenario)
}
