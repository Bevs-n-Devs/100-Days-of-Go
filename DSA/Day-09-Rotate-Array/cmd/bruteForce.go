package main

/*
In the brute force solution, we repeatedly rotate the array one step to the right, k times.

Steps:
1. Create a helper function to rotate the array by 1 step.
2. Perform the single-step rotation k times.
*/

// helper function to rotate the array by one step
func rotateOnce(nums []int) {
	last := nums[len(nums)-1]

	// shift elements to the right
	for i := len(nums) - 1; i > 0; i-- {
		nums[i] = nums[i-1]
	}

	// move the last element to the first position
	nums[0] = last
}

func bruteForceRoateArray(nums []int, k int) []int {
	// brute force
	n := len(nums)
	k = k % n // handle cases where k >= n

	// rotate the array k times
	for i := 0; i < k; i++ {
		rotateOnce(nums)
	}

	return nums
}
