package main

import (
	"fmt"
)

// Pointers are more efficient when working with large data structures since only the memory address is passed, not the entire object.

type LargeStruct struct {
	data [1000]int // list of 1000 integers
}

func process(ks *LargeStruct) {
	// modify the original data
	ks.data[0] = 42
}

func efficientMemory() {
	fmt.Println("\nEfficient Memory Usage with Pointers")

	info :=
		`- Pointers are more efficient when working with large data structures since only the memory address is passed, not the entire object.
	`
	fmt.Println(info)

	ls := LargeStruct{}
	process(&ls) // passing to a pointer
	fmt.Println("Fisrt element:", ls.data[0])
}
