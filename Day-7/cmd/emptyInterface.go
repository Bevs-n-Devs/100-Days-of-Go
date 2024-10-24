package main

import "fmt"

/*
The empty interface (interface{}) is special because it can represent any type.
Since every type in Go implements at least zero methods, all types satisfy the empty interface.

This is useful when you need to work with different types,
but it sacrifices some type safety because you won’t know the specific type at compile time.
*/

func printAny(value interface{}) {
	fmt.Println(value)
}

func emptyInterface() {
	fmt.Println("\nEmpty Interface")

	printAny(42)
	printAny("Hello world, hello Yaw!")
	printAny(3.14)
}
