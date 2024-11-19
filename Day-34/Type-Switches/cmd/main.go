package main

import "fmt"

/*
Type switches allow you to handle multiple possible types that an interface can hold.
Instead of asserting a single type, you match against several types in a clean and readable way.
*/

func main() {
	fmt.Println("hello world, hello Yaw!")

	fmt.Println("\nType switches")

	typeSwitchWithBasicTypes()

	typeSwitchWithStructs()
}
