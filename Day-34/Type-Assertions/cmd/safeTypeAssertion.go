package main

import "fmt"

/*
Best Practice:
- Always use the ok idiom unless you are absolutely sure of the type.
*/

func safeTypeAssertions() {
	fmt.Print("\nSafe Type Assertions")

	// create an empty interface holding a string
	var data interface{} = "Learning Go!"

	// safe type assertion
	str, ok := data.(string) // checks if the data object is a string with .(string) -> returns bool
	if ok {
		fmt.Println("\nExtracted string:", str)
	} else {
		fmt.Println("Data is not a string.")
	}

	// unsafe type assertion
	fmt.Println("\nType Assertion Panics:")
	fmt.Println("If the assertion fails, a panic will be rasied. Below I am trying to check if a string is an integer.")
	num := data.(int) // will panic if the assertion fails
	fmt.Println(num)
}
