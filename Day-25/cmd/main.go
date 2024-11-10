package main

import (
	"fmt"
)

/*
Reflection is a feature in Go that allows you to inspect the type and value of variables at runtime.
This can be particularly useful when you need to write functions or libraries that can operate on a variety of types without knowing the types in advance.

In Go, the reflect package provides the functionality needed for reflection. Here are the primary elements in the reflect package:

1. reflect.Type: Represents the type of a Go value.
2. reflect.Value: Holds the actual value, which can be inspected and modified.
*/

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nExploring Reflection in Go")

	reflectionBasics()

	accessingFieldsWithReflection()

	modifyingValuesWithReflect()
}
