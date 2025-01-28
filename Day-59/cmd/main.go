package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println("Hello wold, hello Yaw!")
	fmt.Println("\nWriting Idiomatic Go Code")

	var whatAreIdioms = `
Go idioms are patternms, practices, and conventions widely adopted by the Go community to write clean, 
efficient, and maintainable code. Writing idomatic Go means leveraging the language's unique design and
strengths, like simplicity, clarity, and performance, rather than tryign to force patterns from other
programming languages.
	`
	fmt.Println(whatAreIdioms)

	var whyIdioms = `
	Why Focus on Idiomatic Go?
		1. Improves Readability: Your code will feel natural to other Go developers.
		2. Reduces Cognitive Load: Idiomatic Go is simple and avoids unnecessary abstraction.
		3. Promotes Maintainability: Teams can work on your code more efficiently.
		4. Leverages Language Strengths: Go's features are optimized for certain patterns (e.g., concurrency, standard library).
	`
	fmt.Println(whyIdioms)

	fmt.Println("\nKey Principles of Idiomatic Go")

	Example1()
	Example2()
	Example3()
	Example4()
	Example5()
	Example6()
	Example7()
	Example8()
	Example9()
	Example10()

	fmt.Println("\nCommon Go Idioms")

	var errorHandling = `
	Error Handling Idiom
	if err != nil {
		return fmt.Errorf("operation failed: %w", err)
	}
	`
	fmt.Println(errorHandling)

	var recieverMethods = `
	Reciever Methods
		- Use value receivers for small structs and pointer receivers for large structs or when modifying the receiver.
	
	Code Example:
	type User struct {
		Name string
	}

	func (u User) Greet() string { // Value receiver
		return "Hello, " + u.Name
	}

	func (u *User) UpdateName(newName string) { // Pointer receiver
		u.Name = newName
	}
	`
	fmt.Println(recieverMethods)

	var summary = `
	Summary Checklist for Writing Idiomatic Go
		- Keep your code simple and readable.
		- Follow Go conventions (gofmt, naming, etc.).
		- Use built-in features like slices, maps, and defer.
		- Error handling should be explicit and close to the source.
		- Write small, focused functions.
		- Favor composition over inheritance.
		- Use interfaces to define behavior, not data.
		- Leverage goroutines and channels for concurrency.

By adopting these idioms, you ensure that your Go code aligns with community standards, making it clean, 
maintainable, and highly effective for solving real-world problems.
	`
	fmt.Println(summary)
}
