package main

import (
	"fmt"
)

/*
With generics, you can also create helper functions for slices, maps, and other data types.

Here’s a generic function to filter elements in a slice based on a condition
*/

// Filter function that takes a slice of T and a predicate function
func Filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func filteringSlice() {
	fmt.Println("\nGenerics in Slices & Maps: Filtering A Slice")

	numbers := []int{1, 2, 3, 4, 5}
	evenNumbers := Filter(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println(evenNumbers) // [2, 4]

	strings := []string{"apple", "banna", "avocado"}
	aStrings := Filter(strings, func(s string) bool {
		return s[0] == 'a'
	})
	fmt.Println(aStrings) // [apple, avocado]
}
