package main

func bruteForceMoveZero(nums []int) []int {
	// create a temp arrary to store non-zero elements
	var tempArray []int

	// count the number of zeros in the array
	zeroCount := 0

	// loop through the array and separate non-zeros
	for _, num := range nums {
		if num != 0 {
			tempArray = append(tempArray, num)
		} else {
			zeroCount++
		}
	}

	// rebuild the original array
	// first place all non-zero elements
	for i := 0; i < len(tempArray); i++ {
		nums[i] = tempArray[i]
	}

	// then place all zeros at the end
	for i := len(tempArray); i < len(nums); i++ {
		nums[i] = 0
	}

	return nums
}
