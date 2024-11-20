package main

/*
The optimal in-place solution involves reversing sections of the array.

Steps:
1. Reverse the entire array.
2. Reverse the first k elements.
3. Reverse the remaining n - k elements.
*/

// helper function to reverse a section of the array
func reverse(nums []int, start int, end int) {
	// swap elements from start to end
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func reverseRotateArray(nums []int, k int) []int {
	// reverse algorithm solution
	n := len(nums)
	k = k % n // handle cases where k >= n

	// step 1: reverse the entire array
	reverse(nums, 0, n-1)

	// step 2: reverse the first k elements
	reverse(nums, 0, k-1)

	// step 3: reverse the remaining n - k elements
	reverse(nums, k, n-1)

	return nums
}
