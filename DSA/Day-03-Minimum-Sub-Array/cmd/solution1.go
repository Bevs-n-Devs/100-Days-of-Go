package main

func solution1(nums []int) int {

	// initialise the number as the first list item
	result := nums[0]

	// start total at 0 to keep track of total sum
	total := 0

	// loop through elements of the list
	for _, num := range nums {
		// add the current numbert to the total
		total += num

		// update result if the total is greater than current result
		if total > result {
			result = total
		}

		// if total becomes negative, reset it to 0 to start new array
		if total < 0 {
			total = 0
		}
	}
	return result
}
