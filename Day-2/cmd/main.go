package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world, hello Yaw!")
	fmt.Println("Basic Go Syntax")

	// variable declaration and initialization
	myName := "Yaw"
	myAge := 38
	myHeight := 5.11
	isDeveloper := true

	// declare contants (a value that doesn't change)
	const pi = 3.14159
	const course = "Golang Basics"

	// type conversion
	var score int = 89
	var percentage float64 = float64(score) / 100.0

	fmt.Println("Course:", course)
	fmt.Println("Name:", myName)
	fmt.Println("Age:", myAge)
	fmt.Println("Height:", myHeight)
	fmt.Println("Is a Developer:", isDeveloper)
	fmt.Println("Pi:", pi)
	fmt.Println("Score as percentage:", percentage)

	fmt.Println("Getting Usrer Input")
	fmt.Println("What is your name: ")

	// get the value of the string object via its memorty allocation
	var name string
	fmt.Scanln(&name)

	// print back the returned value
	fmt.Printf("Your name is %v", name)

	// call the exercies function from exercises.go
	fmt.Println("\n\nExercise 1:")
	exercises()

	fmt.Println("\nExercise 2:")
	exercise2()

}
