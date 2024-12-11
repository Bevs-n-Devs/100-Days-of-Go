package main

import "fmt"

func stackNextGreaterElement(nums1 []int, nums2 []int) []int {
	fmt.Println("\nStack Solution")

	// Map to store the next greater element for each number in nums2
	nextGreaterMap := make(map[int]int)

	// Monotonic decreasing stack to process nums2
	stack := []int{}

	// Traverse nums2 from right to left
	for i := len(nums2) - 1; i >= 0; i-- {
		num := nums2[i]

		// Maintain the stack: pop elements smaller than or equal to the current num
		for len(stack) > 0 && stack[len(stack)-1] <= num {
			stack = stack[:len(stack)-1]
		}

		// If stack is empty, there is no next greater element
		if len(stack) == 0 {
			nextGreaterMap[num] = -1
		} else {
			nextGreaterMap[num] = stack[len(stack)-1]
		}

		// Push the current number onto the stack
		stack = append(stack, num)
	}

	// Use the precomputed map to find results for nums1
	result := make([]int, len(nums1))
	for i, num := range nums1 {
		result[i] = nextGreaterMap[num]
	}

	return result
}

func stackSolution() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	result := stackNextGreaterElement(nums1, nums2)
	fmt.Println(result)

	num3 := []int{2, 4}
	num4 := []int{1, 2, 3, 4}
	result2 := stackNextGreaterElement(num3, num4)
	fmt.Println(result2)
}
