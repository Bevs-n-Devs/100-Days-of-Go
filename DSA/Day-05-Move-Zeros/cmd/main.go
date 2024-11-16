package main

import "fmt"

/*
Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.

The problem asks you to move all the zeros in an array to the end while maintaining the relative order of the non-zero elements.
After moving the zeros, you should not return anything — the array should be modified in place.

Example 1:
Input: nums = [0, 1, 0, 3, 12]
Output: [1, 3, 12, 0, 0]

Example 2:
Input: nums = [0, 0, 1]
Output: [1, 0, 0]
*/

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nDay 5 of DSA: Move Zeros - Leetcode #283")

	fmt.Println("\nSolution 1: Two Pointers Solution")

	case1 := []int{0, 1, 0, 3, 12}
	case2 := []int{0, 0, 1}

	fmt.Printf("Example 1: %v\nOutput: %v\n", case1, moveZeroes(case1))
	fmt.Printf("Example 2: %v\nOutput: %v\n", case2, moveZeroes(case2))

	fmt.Println("\nSolution2: Brute Force Solution")

	fmt.Printf("Example 1: %v\nOutput: %v\n", case1, bruteForceMoveZero(case1))
	fmt.Printf("Example 2: %v\nOutput: %v\n", case2, bruteForceMoveZero(case2))
}
