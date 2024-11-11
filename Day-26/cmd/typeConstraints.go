package main

import (
	"fmt"
)

/*
Type constraints restrict the types that can be passed to a generic function or type. Go provides several constraints out of the box:

1. any: Allows any type.
2. ~: Indicates that a type is compatible with another type (e.g., ~int for custom integer types).
3. comparable: Limits the type to types that can be compared with == and !=.

You can also define custom type constraints as well.
*/

// any constraints example

// PrintElements accepts a slice of any type and prints each element
func PrintElements[T any](elements []T) {
	for _, element := range elements {
		fmt.Println(element)
	}
}

/*
The any constraint allows the type parameter to accept any data type. This is the most flexible constraint and is
useful when you just need to work with multiple types without any specific operations on them.

Here, PrintElements takes a slice of any type T. The any constraint allows us to pass slices of int, string, float64, etc., and the function simply prints each element.
*/

// ~ (Tilda) constraint example

// define a custom constraint for integer-like types
type Integerish interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Double function that doubl;es the value of any interger-based type
func Double[T Integerish](value T) T {
	return value * 2
}

// define a custom integer type
type MyInt int

/*
The ~ constraint lets you specify that a type must be based on another type. This is particularly useful when
working with custom types that should act like underlying types (e.g., type Celsius float64), and you want to
allow both custom types and the underlying type in your generic function.

Here, Integerish uses ~ to specify types that are based on int, such as int, int8, MyInt, etc. This allows
Double to accept both standard integer types and custom integer types like MyInt.
*/

// comparable constraint example

// Contains checks if a slice contains a specific value
func Contains[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

/*
The comparable constraint allows the type parameter to work only with types that support equality comparison (== and !=).
This is useful for operations that involve comparisons, such as checking if an item exists in a slice.

Here, Contains checks if a given value exists in a slice. The comparable constraint ensures that only types
that support == and != (e.g., int, string) are allowed, making it safe to use == within the function.
*/

func typeConstraints() {
	fmt.Println("\nType Constraints")

	fmt.Println("\nPrint any Type Elements")
	PrintElements([]int{1, 2, 3})
	PrintElements([]string{"a", "b", "c"})
	PrintElements([]float64{1.1, 2.2, 3.3})

	fmt.Println("\nPrint Constraints Using Tilda ~")
	fmt.Println(Double(5))
	fmt.Println(Double(MyInt(3)))

	fmt.Println("\nComparable Constraints Example")
	fmt.Println(Contains([]int{1, 2, 3}, 2))
	fmt.Println(Contains([]string{"a", "b", "c"}, "d"))
}
