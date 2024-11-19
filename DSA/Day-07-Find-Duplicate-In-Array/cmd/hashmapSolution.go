package main

// Using a hash map (or a Go map), we can efficiently count the frequency of each element and find duplicates.

func findDuplicatesHashmap(nums []int) []int {
	// hashmap solution
	hashMap := make(map[int]int) // map to count the elements
	duplicates := []int{}

	// count how many times each element appears
	for _, num := range nums {
		hashMap[num]++
	}

	// collect the duplicates
	for num, count := range hashMap {
		if count > 1 {
			duplicates = append(duplicates, num)
		}
	}

	return duplicates
}
