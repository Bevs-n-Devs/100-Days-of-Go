package main

import "fmt"

func dailyTempStack(temperatures []int) []int {
	// stack solution
	fmt.Println("\nStack Solution")

	n := len(temperatures)
	answer := make([]int, n)
	stack := []int{} // Stack to store indices

	for i := 0; i < n; i++ {
		// While stack is not empty and the current temperature is greater than the temperature at the index on top of the stack
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // Pop the top of the stack
			answer[top] = i - top        // Calculate the number of days
		}
		stack = append(stack, i) // Push current index onto the stack
	}

	return answer
}

func stackSolution() {
	case1 := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Printf("Example 1: %v\nOutput: %v\n", case1, dailyTempStack(case1))

	case2 := []int{30, 40, 50, 60}
	fmt.Printf("Example 2: %v\nOutput: %v\n", case2, dailyTempStack(case2))

	case3 := []int{30, 60, 90}
	fmt.Printf("Example 3: %v\nOutput: %v\n", case3, dailyTempStack(case3))
}
