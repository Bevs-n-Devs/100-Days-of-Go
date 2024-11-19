package main

/*
The brute force approach checks all possible pairs of lines to calculate the area of water they can hold.
For each pair, the area is the product of the shorter line's height and the distance between the two lines.

Steps:
1. Use two nested loops to iterate through all pairs of lines.
2. For each pair, calculate the area as min(height[i], height[j]) * (j - i).
3. Keep track of the maximum area encountered.
*/

func bruteForceContainerWithMostWater(height []int) int {
	// brute force solution
	maxArea := 0

	// loop through all pairs of lines
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			// calculate the area for the pair of lines
			h := min(height[i], height[j])
			width := j - i
			area := h * width

			// update maxArea if the current area is larger
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}
