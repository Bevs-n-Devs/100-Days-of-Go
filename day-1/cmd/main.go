package main

import "fmt"

func main() {
	fmt.Println("Hello world, hello Yaw!")
	fmt.Println("Day 1 of 100 Days of Go: Contains Duplicate Algorithm")

	var result = containsDuplicate([]int{1, 2, 3, 3, 4})
	fmt.Println(result)
}

func containsDuplicate(nums []int) bool {
	// create an mapped int,bool object (INT index, BOOL value)
	checkDuplicate := make(map[int]bool)

	// loop through the nums int array
	for _, num := range nums {
		// take the value of each num in in nums
		// and pass through checkDuplicate
		if checkDuplicate[num] {
			return true
		}
		// if no duplicate found then mark the num in containsDuplicate as true
		checkDuplicate[num] = true
	}
	// return false if we have looped through all the values and found no duplicate
	// this means all the values from the array have been assigned in containsDuplicate WITHOUT finding a duplicate
	return false
}
