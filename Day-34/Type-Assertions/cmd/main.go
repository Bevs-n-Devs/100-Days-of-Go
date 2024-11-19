package main

import "fmt"

/*
Type assertions are used to retrieve the concrete type from an interface.
An interface in Go can hold any value that implements its methods.
Type assertions allow you to safely extract and work with the concrete type.
*/

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("Type Assertions")

	typeAssertionCustomStruct()

	safeTypeAssertions()

}
