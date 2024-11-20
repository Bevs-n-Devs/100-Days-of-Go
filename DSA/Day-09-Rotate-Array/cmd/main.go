package main

import "fmt"

/*
Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.

Basically have to rotate the array k times.

Example:
Input: nums = [1,2,3,4,5,6,7], k = 3
Output: [5,6,7,1,2,3,4]

Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]
*/

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nDay 9 of DSA: Rotate Array - Leetcode #189")

	case1 := []int{1, 2, 3, 4, 5, 6, 7}
	k1 := 3
	case2 := []int{-1, -100, 3, 99}
	k2 := 2

	fmt.Println("\nSolution 1: Brute Force Solution")
	fmt.Printf("Example 1: %v, k: %d\nOutput: %v\n", case1, k1, bruteForceRoateArray(case1, k1))
	fmt.Printf("Example 2: %v, k: %d\nOutput: %v\n", case2, k2, bruteForceRoateArray(case2, k2))

	fmt.Println("\nSolution 2: Hash Map Solution")
	fmt.Printf("Example 1: %v, k: %d\nOutput: %v\n", case1, k1, hashMapRotateArray(case1, k1))
	fmt.Printf("Example 2: %v, k: %d\nOutput: %v\n", case2, k2, hashMapRotateArray(case2, k2))

	fmt.Println("\nSolution 3: Reverse Solution")
	fmt.Printf("Example 1: %v, k: %d\nOutput: %v\n", case1, k1, reverseRotateArray(case1, k1))
	fmt.Printf("Example 2: %v, k: %d\nOutput: %v\n", case2, k2, reverseRotateArray(case2, k2))
}
