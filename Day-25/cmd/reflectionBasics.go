package main

import (
	"fmt"
	"reflect"
)

/*
reflect.TypeOf() retrieves the type of a variable.
reflect.ValueOf() retrieves the current value.
*/

func reflectionBasics() {
	fmt.Println("\nRelection Basics")

	var number int = 42
	var text string = "hello"

	// get the Type and Value of `number`
	fmt.Println("Type of `number`:", reflect.TypeOf(number))
	fmt.Println("Value of `number`:", reflect.TypeOf(number))

	// get the type & value of `text`
	fmt.Println("Type of `text`:", reflect.TypeOf(text))
	fmt.Println("Value of `text`:", reflect.ValueOf(text))
}
