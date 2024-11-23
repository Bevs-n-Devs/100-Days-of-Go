package main

import "fmt"

func pointersToArrays() {
	fmt.Println("\nPointers to Arrays")

	arr := [3]int{1, 2, 3}
	p := &arr

	fmt.Println("Array through pointer:", (*p)[1]) // access areray element
	(*p)[1] = 42                                   // modify array element @ index 1
	fmt.Println("Modified array:", arr)
}
