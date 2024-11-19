package main

import "math"

/*
If modifying the input array is allowed, we can use the values as indices to mark visited elements.
This approach works only when the array contains integers from 1 to n (inclusive).

1. Iterate through the array and for each number, mark th1. e index corresponding to its value as visited (e.g., by negating the value).
2. If the value at the corresponding index is already negative, it means the number is a duplicate.
*/

func findDuplicatesInPlace(nums []int) []int {
	// in place solution
	duplicates := []int{}

	for _, num := range nums {
		// use the value as an index (absolute value for safety)
		index := int(math.Abs(float64(num))) - 1

		// check if the value of at this index is alread a negative
		if nums[index] < 0 {
			duplicates = append(duplicates, index+1)
		} else {
			// mark the index visited as a negative number
			nums[index] = -nums[index]
		}
	}

	return duplicates
}
