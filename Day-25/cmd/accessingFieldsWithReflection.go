package main

import (
	"fmt"
	"reflect"
)

/*
Reflection can be useful when working with structs, especially when you don’t know the field names or types in advance.

reflect.ValueOf(input) gives the reflect value of the input.
NumField() and Field(i) are used to access each field in the struct dynamically.
*/

type User struct {
	ID   int
	Name string
	Age  int
}

func printFields(input interface{}) {
	val := reflect.ValueOf(input)

	// check if the input is a struct
	if val.Kind() == reflect.Struct {
		for i := 0; i < val.NumField(); i++ {
			field := val.Type().Field(i)
			value := val.Field(i)
			fmt.Printf("%s: %v\n", field.Name, value)
		}
	} else {
		fmt.Println("Expected a struct")
	}
}

func accessingFieldsWithReflection() {
	fmt.Println("\nAccessing Fields With Reflection")

	user := User{ID: 1, Name: "Alice", Age: 25}
	printFields(user)
}
