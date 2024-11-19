package main

// The brute force approach involves comparing every element with every other element. This can be achieved using two nested loops.

func findDuplicatesBruteForce(nums []int) []int {
	// brute force solution
	n := len(nums)
	duplicates := make([]int, 0)

	// outer loop to loop through each element
	for outer := 0; outer < n; outer++ {
		// inner loop to compare each element
		for inner := outer + 1; inner < n; inner++ {
			// check if the current element is equal to the inner element
			if nums[outer] == nums[inner] {
				// check if the current element is already in the duplicates list
				alreadyAdded := false
				for _, val := range duplicates {
					if val == nums[outer] {
						alreadyAdded = true
						break
					}
				}
				if !alreadyAdded {
					duplicates = append(duplicates, nums[outer])
				}
			}
		}
	}

	return duplicates
}
