package main

import (
	"fmt"
	"strings"
)

// twoSum finds two numbers in nums that add up to the target and return their indices
func twoSum(nums []int, target int) []int {
	// create a hash map to store each number and its index
	hashMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		// check if the complement of the current number exists in the hashmap
		if j, found := hashMap[complement]; found {
			// if found, return the indices
			return []int{j, i}
		}
		// store the current number and its index in the map
		hashMap[num] = i
	}
	// if no solution is found, return an empty slice
	return []int{}
}

func twoSumAlgorithm() {
	fmt.Println("\n"+strings.Repeat("*", 4), "Two Sums Alogrithm Challenge", strings.Repeat("*", 4))

	nums := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(nums, target)

	if len(result) > 0 {
		fmt.Printf("Indices: %d, %d\n", result[0], result[1])
	} else {
		fmt.Println("No solution found")
	}
}
