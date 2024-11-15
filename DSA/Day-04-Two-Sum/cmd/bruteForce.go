package main

func twoSumBruteForce(nums []int, target int) []int {
	// check each possible pair by:

	// looping through the current index
	for currentIndex := 0; currentIndex < len(nums); currentIndex++ {
		// looping through the next index
		for nextIndex := currentIndex + 1; nextIndex < len(nums); nextIndex++ {
			// check if the two indexes sum to the target
			if nums[currentIndex]+nums[nextIndex] == target {
				// return indexes of numbers that sum up to the target
				return []int{currentIndex, nextIndex}
			}
		}
	}

	// return nil if no pair found
	return nil
}
