package main

import "fmt"

/*
You can define a custom constraint for more specific use cases.

Numeric is a custom constraint that limits the Sum function to work only with int and float64.
*/

// define a custom constraint for numeric types
type Numneric interface {
	~int | ~float64
}

// Sum function that works with any numeric type
func Sum[T Numneric](a, b T) T {
	return a + b
}

func typeSetsCustomConstraints() {
	fmt.Println("\nType Sets: Custom Constraints")

	fmt.Println(Sum(10, 20))     // works with integers
	fmt.Println(Sum(3.14, 2.71)) // works with floats
}
