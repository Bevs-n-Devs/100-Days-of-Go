package main

import "fmt"

/*
Create a program that asks for the user's name, age, and favorite number.
Convert the favorite number to a float and then multiply it by 3.14 (simulating Pi).
Print the result back to the user.
*/

func exercises() {
	fmt.Println("Go Exercises: Basic Syntax and Getting User Input")

	// declare varaibles for user input
	var userName string
	var userAge int
	var favouriteNumber int

	// use fmt.Scanln() to get the user input
	fmt.Println("What is your name?")
	fmt.Scanln(&userName)

	fmt.Println("What is your age?")
	fmt.Scanln(&userAge)

	fmt.Println("What is your favourite number?")
	fmt.Scanln(&favouriteNumber)

	// Multiply the favourite number by Pi (3.14)
	const pi = 3.14
	new_value := float64(favouriteNumber) * pi

	// print the result
	fmt.Printf("\nHello %v! Your age is %v and your favourite number is %v", userName, userAge, favouriteNumber)
	fmt.Printf("\nYour favourite number %v multiplied by Pi: %v", favouriteNumber, new_value)
}
