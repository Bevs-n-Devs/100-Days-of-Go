package main

import (
	"fmt"
	"reflect"
)

/*
Reflection also allows for modification of values, but you can only modify values if they are settable.
To make a value settable, it needs to be passed as a pointer.

FieldByName("Age") finds the Age field by its name.
CanSet() checks if the field can be set, ensuring it’s not a private field or constant.
SetInt() sets the integer value of the field.
*/

func modifyAge(input interface{}) {
	val := reflect.ValueOf(input)

	// make sure input is a pointer to a struct
	if val.Kind() == reflect.Ptr && val.Elem().Kind() == reflect.Struct {
		val = val.Elem() // dereference the pointer
		ageField := val.FieldByName("Age")

		// check if the Age field can be set
		if ageField.IsValid() && ageField.CanSet() && ageField.Kind() == reflect.Int {
			ageField.SetInt(35) // modify the age field
		}
	} else {
		fmt.Println("Expected a pointer to a struct")
	}
}

func modifyingValuesWithReflect() {
	fmt.Println("\nModifying Values With Reflection")

	user := &User{ID: 1, Name: "Alice", Age: 25}
	modifyAge(user)
	fmt.Println("Modified user:", *user)
}
