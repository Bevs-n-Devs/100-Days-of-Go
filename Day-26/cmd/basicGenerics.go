package main

import "fmt"

/*
Type parameters are defined within square brackets ([]) right after the function or type name.

T is a type parameter that can represent any type (using the any constraint).
The function PrintValue can accept any type T for the parameter value.
*/

// generic function that accepts any type T
func PrintValue[T any](value T) {
	fmt.Println(value)
}

func typeParameters() {
	fmt.Println("\nBasics of Type Parameters")

	PrintValue(42)      // using an int
	PrintValue("hello") // using a string
	PrintValue(3.14)    // usign a float
}
