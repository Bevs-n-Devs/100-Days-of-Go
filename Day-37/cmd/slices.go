package main

import "fmt"

// helper function to modify a slice
func modifySlice(s []int) {
	// modify the original slice
	s[0] = 42
}

func slices() {
	fmt.Println("\nSlices")

	info :=
		`- Slices internally use pointers for their underlying array.
 `
	fmt.Println(info)

	arr := []int{1, 2, 3}
	fmt.Println("Original array:", arr)

	modifySlice(arr)
	fmt.Println("Modified array:", arr) // update the first element
}
