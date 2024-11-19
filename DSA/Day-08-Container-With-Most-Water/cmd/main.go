package main

import "fmt"

/*
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

Example
Input: height = [1,8,6,2,5,4,8,3,7]
Output
*/
func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Print("\nDay 8 of DSA: Container With Most Water - Leetcode #11")

	case1 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	case2 := []int{1, 1}

	fmt.Println("\nSolution 1: Brute Force Solution")
	fmt.Printf("Example 1: %v\nOutput: %d\n", case1, bruteForceContainerWithMostWater(case1))
	fmt.Printf("Example 2: %v\nOutput: %d\n", case2, bruteForceContainerWithMostWater(case2))

	fmt.Println("\nSolution 2: Two Pointers Solution")
	fmt.Printf("Example 1: %v\nOutput: %d\n", case1, twoPointerContainerWithMostWater(case1))
	fmt.Printf("Example 2: %v\nOutput: %d\n", case2, twoPointerContainerWithMostWater(case2))
}
