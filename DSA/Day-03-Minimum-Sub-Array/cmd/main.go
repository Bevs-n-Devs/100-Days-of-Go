package main

import "fmt"

/*
Given an integer array nums, find the contiguous subarray (containing at least one number) that has the largest sum, and return that sum.

Example:
Input: nums = [-2, 1, -3, 4, -1, 2, 1, -5, 4]
Output: 6
Explanation: The subarray [4, -1, 2, 1] has the largest sum = 6.
*/

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nDay 3 of DSA: Minimum Sub Array - Leetcode #53")

	case1 := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	case2 := []int{1}
	case3 := []int{5, 4, -1, 7, 8}

	fmt.Println("\nConverted Python Solution")
	fmt.Printf("Example 1: %v\nOutput: %d\n", case1, solution1(case1))
	fmt.Printf("Example 2: %v\nOutput: %d\n", case2, solution1(case2))
	fmt.Printf("Example 3: %v\nOutput: %d\n", case3, solution1(case3))

	fmt.Println("\nSolution 2: Kadane's Algorithm using Math Package")
	fmt.Printf("Example 1: %v\nOutput: %d\n", case1, maxSubArray(case1))
	fmt.Printf("Example 2: %v\nOutput: %d\n", case2, maxSubArray(case2))
	fmt.Printf("Example 3: %v\nOutput: %d\n", case3, maxSubArray(case3))
}
