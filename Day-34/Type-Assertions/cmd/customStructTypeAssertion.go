package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func displayData(data interface{}) {
	if person, ok := data.(Person); ok {
		fmt.Printf("Person Name: %s, Age: %d\n", person.Name, person.Age)
	} else {
		fmt.Println("Not a person")
	}
}

func typeAssertionCustomStruct() {
	fmt.Println("\nType Assertions - Custom Struct")

	p := Person{"John Doe", 30}
	displayData(p)
	displayData("Just a string!")
}
