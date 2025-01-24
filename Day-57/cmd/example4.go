package main

import "fmt"

func Example4() {
	fmt.Println("\nExample 4: User Interfaces")
	var scenario = `
	Violation of PoLA:
		- A button labeled "Cancel" behaves like a "Save" button.
		- Users expect a "Cancel" button to discard changes, not save them.

	Following PoLA:
		- Label the button "Save" if it performs a save action, or separate the "Save" and "Cancel" actions entirely.
	`
	fmt.Println(scenario)
}
