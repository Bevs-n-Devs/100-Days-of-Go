package main

import (
	"fmt"
)

/*
Given an integer array nums of length n where all the integers of nums are in the range [1, n] and each
integer appears at most twice, return an array of all the integers that appears twice.

You must write an algorithm that runs in O(n) time and uses only constant auxiliary space, excluding the space needed to store the output

Finding duplicates in an array involves identifying elements that appear more than once.
There are multiple ways to approach this problem, from brute force to hashmap solutions.

Example
Input: nums = [4, 3, 2, 7, 8, 2, 3, 1]
Output: [2, 3]
*/

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nDay 7 of DSA: Find Duplicate Numbers - Leetcode #442")

	case1 := []int{4, 3, 2, 7, 8, 2, 3, 1}
	case2 := []int{3, 1, 3, 4, 2}
	case3 := []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}

	fmt.Println("\nSolution 1: Brute Force Solution")
	fmt.Printf("Example 1: %v\nOutput: %v\n", case1, findDuplicatesBruteForce(case1))
	fmt.Printf("Example 2: %v\nOutput: %v\n", case2, findDuplicatesBruteForce(case2))
	fmt.Printf("Example 3: %v\nOutput: %v\n", case3, findDuplicatesBruteForce(case3))

	fmt.Println("\nSolution 2: Hashmap Solution")
	fmt.Printf("Example 1: %v\nOutput: %v\n", case1, findDuplicatesHashmap(case1))
	fmt.Printf("Example 2: %v\nOutput: %v\n", case2, findDuplicatesHashmap(case2))
	fmt.Printf("Example 3: %v\nOutput: %v\n", case3, findDuplicatesHashmap(case3))

	fmt.Println("\nSolution 3: In Place Solution")
	fmt.Printf("Example 1: %v\nOutput: %v\n", case1, findDuplicatesInPlace(case1))
	fmt.Printf("Example 2: %v\nOutput: %v\n", case2, findDuplicatesInPlace(case2))
	fmt.Printf("Example 3: %v\nOutput: %v\n", case3, findDuplicatesInPlace(case3))
}
