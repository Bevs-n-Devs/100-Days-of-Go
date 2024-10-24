package main

import "fmt"

/*
When working with the empty interface (interface{}), you may need to access the original type.
Type assertions are used to retrieve the dynamic value stored in an interface.

The line str, ok := value.(string) checks if value is of type string.
If it is, the assertion succeeds, and the ok variable is true.
Otherwise, ok is false, and you can handle the case accordingly.
*/

func printType(value interface{}) {
	// type assertion
	str, ok := value.(string)
	if ok {
		fmt.Println("It's a string:", str)
	} else {
		fmt.Println("This is not a string")
	}
}

func typeAssertions() {
	fmt.Println("\nType Assertions with Interfaces")

	printType("Hello")
	printType(1986)
}
