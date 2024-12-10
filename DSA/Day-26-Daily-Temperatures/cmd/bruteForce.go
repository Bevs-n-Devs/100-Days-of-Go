package main

import "fmt"

/*
Approach:

1. For each day, look at all future days to find the first day with a higher temperature.
2. Calculate the difference in days.

Time Complexity: O(n^2), where n is the number of days
Space Complexity: O(1), excluding the output array
*/

func dailyTempBruteForce(temperatures []int) []int {
	// brute force
	fmt.Println("\nBrute Force Solution")

	n := len(temperatures)
	answer := make([]int, n)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if temperatures[j] > temperatures[i] {
				answer[i] = j - i
				break
			}
		}
	}

	return answer
}

func bruteFoceSolution() {
	case1 := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Printf("Example 1: %v\nOutput: %v\n", case1, dailyTempBruteForce(case1))

	case2 := []int{30, 40, 50, 60}
	fmt.Printf("Example 2: %v\nOutput: %v\n", case2, dailyTempBruteForce(case2))

	case3 := []int{30, 60, 90}
	fmt.Printf("Example 3: %v\nOutput: %v\n", case3, dailyTempBruteForce(case3))
}
