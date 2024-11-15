package main

func twoSumHashMap(nums []int, target int) []int {
	// create a map to store each index & number
	numMap := make(map[int]int)

	for index, num := range nums {
		// calculate the result
		result := target - num

		// chekc if the result is in map
		if idx, found := numMap[result]; found {
			// if found return indexes
			return []int{idx, index}
		}

		// store the current number with the index in the map
		numMap[num] = index
	}

	// return nil if no pair found
	return nil
}
