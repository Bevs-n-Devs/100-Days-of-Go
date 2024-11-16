package main

func moveZeroes(nums []int) []int {
	// pointer for the position of th enext non-zero element
	nextNonZero := 0

	// loop through the array
	for i := 0; i < len(nums); i++ {
		// if the current element is not zero
		if nums[i] != 0 {
			// swap the current element with the element at the next non-zero position
			nums[nextNonZero], nums[i] = nums[i], nums[nextNonZero]
			// increment the next non-zero position
			nextNonZero++
		}
	}

	return nums
}
