package main

import (
	"fmt"
	"math"
)

/*
This algorithm efficiently calculates the maximum sum of a contiguous subarray within a one-dimensional array of numbers in
𝑂(𝑛) time complexity.
*/

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]
	currentSum := nums[0]

	for i := 1; i < len(nums); i++ {
		// update currentSum to either the current element or the sum of current element and currentSum
		currentSum = int(math.Max(float64(nums[i]), float64(currentSum+nums[i])))
		// update maxSum if currentSum is larger
		maxSum = int(math.Max(float64(maxSum), float64(currentSum)))
	}

	return maxSum
}

func maxSubArrayAlgorithim() {
	fmt.Println("\nKadane's Algorithm: Maximum Sum of A Subarray")
	// example array
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	// finding & printing the maximum sum of subarray
	fmt.Println("The maximum sum of a subarray is:", maxSubArray(nums))
}
