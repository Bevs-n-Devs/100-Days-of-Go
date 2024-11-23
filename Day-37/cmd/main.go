package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nDay 37 of 100 Days of Go: Go Pointers and Structs")

	intro := `
	Introduction

Pointers are a cornerstone of programming, offering a powerful way to work with memory directly. 
This lesson provides a deep dive into pointers, covering their mechanics, use cases, and best practices.
	`
	fmt.Println(intro)

	whatArePointers := `
	What Are Pointers?

A pointer is a variable that holds the memory address of another variable. By using pointers, you can:

- Share large data structures efficiently without copying them.
- Modify data in place (important in functions).
- Work with dynamic data structures (e.g., linked lists, trees).
	`
	fmt.Println(whatArePointers)

	basicTerms := `
	Basic Terms

- Address (&): The memory location of a variable.
- Dereference (*): Accessing the value at a memory location
	`
	fmt.Println(basicTerms)

	// address of operator (&)
	num := 42
	p := &num // points to the memory address of num

	fmt.Println("Value of num:", num)
	fmt.Println("Address of num:", &num) // fetches address where variable num is stored
	fmt.Println("Value of p:", p)        // points to the memory address of num

	// derefence of operator (*)
	fmt.Println("\nValue at address p:", *p) // dereference p to get the value of num
	*p = 100                                 // change the value of num through the pointer
	fmt.Println("Updated value of num:", num)

	pointerUseCases()

	zeroValue()

	pointerToPointer()

	arraysAndSlices()

	bestPractices := `
	BEST PRACTICES

1. Minimize Pointer Usage When Not Needed. Use pointers only when:
	- Modifying the original value.
	- Avoiding large memory copies.
2. Check for nil Before Dereferencing Always ensure pointers are non-nil before dereferencing.
3. Use Struct References for Efficiency When passing structs to functions, prefer using pointers to avoid unnecessary copying.
4. Avoid Excessive Pointer Nesting Pointers to pointers can make code hard to read and maintain.
5. Prefer Slices Over Arrays with Pointers Slices are naturally efficient due to their pointer-like nature.
	`
	fmt.Println(bestPractices)

	commonPitfals := `
	Common Pitfalls

- Uninitialized Pointer Access: Accessing a nil pointer causes runtime panic.
- Dangling Pointers: Using a pointer after the object it references has gone out of scope.
- Confusion with Dereferencing: Be mindful of when to dereference (*) vs. assigning a pointer.
	`
	fmt.Println(commonPitfals)

	conclusion := `
	CONCLUSION

Pointers in Go are a robust tool for efficient memory management and enabling direct data manipulation. 
While they provide power and flexibility, misuse can lead to bugs or performance issues. 
By understanding the & (address of) and * (dereference) operators deeply and practicing best practices, 
you can master pointers and apply them effectively in your Go programs.
	`
	fmt.Print(conclusion)
}
