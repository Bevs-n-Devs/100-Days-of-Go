package main

import "fmt"

func arraysAndSlices() {
	fmt.Println("\nArrays and Slices with Pointers")

	info :=
		`- Pointers can be used to pass arrays and slices as function arguments. This is useful when you need to modify the values in the original array or slice.
	`
	fmt.Println(info)

	pointersToArrays()

	slices()
}
