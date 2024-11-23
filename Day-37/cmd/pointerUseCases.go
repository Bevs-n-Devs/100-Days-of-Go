package main

import "fmt"

func increnment(x *int) {
	// modify the value of x at the address
	*x = *x + 1
}

func pointerUseCases() {
	fmt.Println("\nPointer Use Cases")

	fmt.Println("\nPassing a Pointer to a Function:")
	x := 5
	fmt.Println("Before increment:", x)
	increnment(&x) // pass the memory address of x variable
	fmt.Println("After increment:", x)

	efficientMemory()

	dynamicDataStructures()
}
