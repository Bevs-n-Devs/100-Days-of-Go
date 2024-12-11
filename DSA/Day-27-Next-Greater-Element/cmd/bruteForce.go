package main

import "fmt"

func bruteForceNextGreaterElement(nums1 []int, nums2 []int) []int {
	fmt.Println("\nBrute Force Solution")
	result := make([]int, len(nums1))

	for i, num := range nums1 {
		found := false
		for j := 0; j < len(nums2); j++ {
			if nums2[j] == num {
				// Start searching for the next greater element from j+1
				for k := j + 1; k < len(nums2); k++ {
					if nums2[k] > num {
						result[i] = nums2[k]
						found = true
						break
					}
				}
				if !found {
					result[i] = -1
				}
				break
			}
		}
	}

	return result
}

func bruteForceSolution() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	result := bruteForceNextGreaterElement(nums1, nums2)
	fmt.Println(result)

	num3 := []int{2, 4}
	num4 := []int{1, 2, 3, 4}
	result2 := bruteForceNextGreaterElement(num3, num4)
	fmt.Println(result2)
}
