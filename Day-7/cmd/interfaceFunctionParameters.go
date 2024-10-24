package main

import "fmt"

/*
Interfaces are often used to pass different types of arguments to functions.
The function only needs to know that the argument implements the required interface.
*/

// define Talker interface
type Talker interface {
	Talk() string
}

// function that accepts any type that satisfies the Talker interface
func announce(t Talker) {
	fmt.Println(t.Talk())
}

type Person3 struct {
	name string
}

func (p Person3) Talk() string {
	return "Hello, I am " + p.name
}

type Dog2 struct {
	breed string
}

func (d Dog2) Talk() string {
	return "BARK! I am a " + d.breed
}

func interfaceFunctionParameters() {
	fmt.Println("\nInterface As Function Parameters")

	person := Person3{name: "Mark"}
	dog := Dog2{breed: "Pitbull"}

	// Both Person3 and Dog2 can be passed to announce
	announce(person)
	announce(dog)
}
