package main

// An interface is a type that defines a set of methods, but it doesn’t provide their implementation.
// Any type that implements all of the methods in an interface can be said to "satisfy" the interface.

import "fmt"

// define Speaker interface
type Greeting interface {
	Greet() string
}

// define Person struct
type Person struct {
	name string
	age  int
	city string
}

// implement the Speak method for Person
func (p Person) Greet() string {
	return "Hello, my name is " + p.name
}

func main() {
	fmt.Println("Hello world, hello Yaw!")
	fmt.Println("\nInterfaces with Go")

	fmt.Println("\nDefining Interfaces")
	person := Person{
		name: "John",
		age:  40,
		city: "Newcastle",
	}

	// Person implements the Speak method - this satisfies the Speaker interface
	talk := person

	// call the Speak method via the interface
	fmt.Println(talk.Greet())

	polymorphism()

	interfaceFunctionParameters()

	emptyInterface()

	typeAssertions()

	multipleInterfaces()

	practicalExample()
}
