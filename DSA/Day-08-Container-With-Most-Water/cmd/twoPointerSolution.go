package main

/*
The two-pointer approach is more efficient. It starts with two pointers at the beginning and end
of the array and calculates the area. Depending on the height of the two lines, it moves the
pointer pointing to the shorter line inward, as a taller line might increase the area.

Steps:
1. Start with two pointers: one at the beginning (left) and one at the end (right) of the array.
2. Calculate the area formed by the lines at the left and right pointers.
3. Move the pointer pointing to the shorter line inward (either increment left or decrement right).
4. Repeat until the pointers meet, keeping track of the maximum area.
*/

func twoPointerContainerWithMostWater(height []int) int {
	// two pointer solution
	left := 0
	right := len(height) - 1
	maxArea := 0

	// loop until the two pointers meet
	for left < right {
		// calculate the area between the two pointers
		h := min(height[left], height[right])
		width := right - left
		area := h * width

		// update maxArea if the current area is larger
		if area > maxArea {
			maxArea = area
		}

		// move the pointer pointing to the shorter line inward
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

// helper function to find the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
