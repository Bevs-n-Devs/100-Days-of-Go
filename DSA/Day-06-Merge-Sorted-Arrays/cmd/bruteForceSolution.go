package main

import (
	"sort"
)

/*
Merge the Arrays:
- Copy the m valid elements of nums1 and all elements of nums2 into a temporary array.

Sort the Temporary Array:
- Use a sorting algorithm to sort the combined array.

Copy Back:
- Copy the sorted array back into nums1.
*/

func bruteForceMergeSortedArrays(nums1 []int, m int, nums2 []int, n int) []int {
	// create a temp array to store merged array#
	temp := make([]int, 0, m+n)

	// add all elements from nums1 (up to m) to temp array
	for i := 0; i < m; i++ {
		temp = append(temp, nums1[i])
	}

	// add all elements from nums2 (up to n) to temp array
	for i := 0; i < n; i++ {
		temp = append(temp, nums2[i])
	}

	// sort the temp array
	sort.Ints(temp)

	// copy the sorted array back to nums1
	for i := 0; i < m+n; i++ {
		nums1[i] = temp[i]
	}

	return nums1
}
