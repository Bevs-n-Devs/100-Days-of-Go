package main

import (
	"fmt"
)

/*
Go introduced support for generics with type parameters in Go 1.18, a major addition that allows functions and
types to be written in a type-agnostic way. Generics enable you to write reusable, type-safe code that can work
across multiple types, avoiding the need for code duplication and reliance on interface{} for polymorphism.

Generics allow you to define a function or a data structure that can operate on types without specifying what
those types are in advance. With generics, you can define a type parameter in your function or type, making the
code adaptable to multiple data types.

This allows for:

1. Type Safety: Ensures that only compatible types are used, avoiding runtime type errors.
2. Code Reusability: Enables writing of more reusable functions and data structures.
3. Performance: Generic code in Go is compiled specifically for the type you use, offering better performance
				compared to using empty interfaces (interface{}).
*/

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nGenerics in Go")

	typeParameters()

	typeConstraints()

	genericsWithDataStructures()

	filteringSlice()

	typeSetsCustomConstraints()
}

/*
BEST PRACTICES

1. Choose Constraints Wisely: Constraints ensure type safety and clarify the types expected, making your code easier to understand.
2. Minimize Overuse: Generics add complexity, so use them only when there’s a clear advantage, like in utility functions or common data structures.
3. Benchmark Performance: Generics are optimized for performance, but if you’re working in a performance-critical application, benchmark the code to ensure it meets your needs.
*/
