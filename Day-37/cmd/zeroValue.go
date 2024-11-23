package main

import "fmt"

func zeroValue() {
	fmt.Println("\nZero Value of a Pointer")

	info := `
- A pointer with no value points to nil.
- Dereferencing a nil pointer causes a runtime panic.
	`
	fmt.Println(info)

	var p *int // pointer declared but not initialised with a value
	fmt.Println("Value of p:", p)
	// uncomment to see the runtime panic
	// fmt.Println("*p:", *p) // will cause a runtime panic
}
