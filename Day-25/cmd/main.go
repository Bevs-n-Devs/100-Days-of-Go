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

/*
Use Cases for Reflection
1. JSON Marshalling/Unmarshalling: Libraries use reflection to parse and build JSON without prior knowledge of the data structure.
2. Dependency Injection: Reflectively finding and assigning dependencies at runtime.
3. Generic Functions: Implementing generic functions before Go introduced native generics.
4. Testing and Mocking: Mocking types and comparing unknown values dynamically.

Best Practices and Warnings
1. Use Carefully: Reflection makes code harder to read and reduces compile-time checks.
2. Performance Overhead: Reflection is slower than regular code, as it requires additional runtime checks.
3. Avoid Excessive Modification: Modifying values with reflection can lead to bugs that are hard to debug.
4. Use Native Generics When Possible: With Go's native support for generics, some reflection use cases can now be handled with type parameters.

Summary
Reflection in Go is a powerful tool for dynamically inspecting and modifying variables, types, and structures at runtime.
However, it comes with trade-offs in readability and performance, so it’s best used sparingly and with clear intent.
This knowledge will enable you to build more flexible and dynamic applications, especially when working with libraries,
frameworks, and tools requiring runtime type handling.
*/
