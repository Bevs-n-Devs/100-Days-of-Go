package main

import "fmt"

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
The task is to find two numbers in the array that add up to the target sum and return their indices.
There may be exactly one solution, and we can assume no duplicate values are needed.

Example
Input: nums = [2, 7, 11, 15], target = 9
Output: [0, 1]
Explanation: nums[0] + nums[1] = 2 + 7 = 9
*/

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nDay 4 of DSA: Two Sum #1")

	case1, target1 := []int{2, 7, 11, 15}, 9
	case2, target2 := []int{3, 2, 4}, 6
	case3, target3 := []int{3, 3}, 6

	fmt.Println("\nBrute Force Solution")
	fmt.Printf("Example: %v, target: %d\nOutput: %d\n", case1, target1, twoSumBruteForce(case1, target1))
	fmt.Printf("Example: %v, target: %d\nOutput: %d\n", case2, target2, twoSumBruteForce(case2, target2))
	fmt.Printf("Example: %v, target: %d\nOutput: %d\n", case3, target3, twoSumBruteForce(case3, target3))

	fmt.Println("\nHashmap Solution")
	fmt.Printf("Example: %v, target: %d\nOutput: %d\n", case1, target1, twoSumHashMap(case1, target1))
	fmt.Printf("Example: %v, target: %d\nOutput: %d\n", case2, target2, twoSumHashMap(case2, target2))
	fmt.Printf("Example: %v, target: %d\nOutput: %d\n", case3, target3, twoSumHashMap(case3, target3))
}
