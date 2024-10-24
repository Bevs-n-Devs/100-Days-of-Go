package main

import "fmt"

/*
Interfaces provide a way to write flexible and decoupled code.
You don’t need to know what specific type you are dealing with; you only need to know that it satisfies the interface.

For example, you might have different types that "speak" in different ways (a Person, a Dog, a Robot),
but as long as they all implement the Speak method, you can use them interchangeably through the Speaker interface.

With polymorphism you can use interfaces to work with different types that share the same behavior.
*/

// define Speaker interface
type Speaker interface {
	Speak() string
}

// define Person struct
type Person2 struct {
	name string
	age  int
	city string
}

// define Dog struct
type Dog struct {
	breed  string
	colour string
	gender string
}

// implement Speak method for Person
func (p Person2) Speak() string {
	return "Hello my name is " + p.name
}

// implement Speak method for Dog
func (d Dog) Speak() string {
	return "WOOF! I am a " + d.colour + " " + d.breed
}

func polymorphism() {
	fmt.Println("\nPolymorphism, Interfaces and Go")

	// create instances of Person & Dog
	person := Person2{
		name: "Yaw",
		age:  38,
		city: "London",
	}
	dog := Dog{
		breed:  "Cocka-Poodle",
		colour: "brown",
		gender: "male",
	}

	// create a slice of Speak34 interfaces
	speakers := []Speaker{person, dog}

	// interate through the slice and call Speak method
	for _, speaker := range speakers {
		fmt.Println(speaker.Speak())
	}

}
