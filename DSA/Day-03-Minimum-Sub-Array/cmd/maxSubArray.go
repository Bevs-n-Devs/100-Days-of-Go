package main

import (
	"math"
)

// maxSubArray find the maximum num of a contiguous subarray in nums using Kadane's Algorithm
func maxSubArray(nums []int) int {
	// hold the sum of subarray found
	// initialise to a low number
	maxSum := math.MinInt32

	currentSum := 0

	// iterate through nums array
	for _, num := range nums {
		// start new subarray @ num or add current number [num] to currentSum
		currentSum = max(num, currentSum+num)

		// update if currentSum is greater than maxSum
		maxSum = max(maxSum, currentSum)
	}
	return maxSum
}
