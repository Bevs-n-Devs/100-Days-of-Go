package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Animal struct {
	Name    string
	Species string
}

func describe(entity interface{}) {
	switch e := entity.(type) {
	case Person:
		fmt.Printf("Person Name: %s, Age: %d\n", e.Name, e.Age)
	case Animal:
		fmt.Printf("Animal Name: %s, Species: %s\n", e.Name, e.Species)
	default:
		fmt.Println("Unknown entity")
	}
}

func typeSwitchWithStructs() {
	fmt.Println("\nType Switches - With Structs")

	yaw := Person{"Yaw", 38}
	pet := Animal{Name: "Tiger", Species: "dog"}

	describe(yaw)
	describe(pet)
	describe(100)
}
