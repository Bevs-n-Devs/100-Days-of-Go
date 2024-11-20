package main

/*
In this approach, a temporary array is used to directly calculate the new positions of elements after k rotations.

Steps:
1. Create a temporary array of size n.
2. For each element in the original array, calculate its new index after k rotations.
3. Copy the temporary array back into the original array.
*/

func hashMapRotateArray(nums []int, k int) []int {
	// hash map solution
	n := len(nums)
	k = k % n // handle cases where k >= n

	// create a temporary array to store the new positions of elements after k rotations
	temp := make([]int, n)

	// calculate the new position for each element
	for i := 0; i < n; i++ {
		// calculate the new index
		newIndex := (i + k) % n
		temp[newIndex] = nums[i]
	}

	// copy the temporary array back into the original array
	copy(nums, temp)

	return nums
}
