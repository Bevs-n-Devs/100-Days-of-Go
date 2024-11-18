package main

import "fmt"

/*
You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n, representing the number of elements in nums1 and nums2 respectively.

Merge nums1 and nums2 into a single array sorted in non-decreasing order.

The final sorted array should not be returned by the function, but instead be stored inside the array nums1.
To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should
be merged, and the last n elements are set to 0 and should be ignored. nums2 has a length of n.

Example 1:

Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
Output: [1,2,2,3,5,6]
Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.
*/

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nMerge Sorted Arrays - Leetcode #88")

	case1_nums1 := []int{1, 2, 3, 0, 0, 0}
	case1_m := 3
	case1_nums2 := []int{2, 5, 6}
	case1_n := 3

	case2_nums1 := []int{1}
	case2_m := 1
	case2_nums2 := []int{}
	case2_n := 0

	case3_nums1 := []int{0}
	case3_m := 0
	case3_nums2 := []int{1}
	case3_n := 1

	fmt.Println("\nMerge Sorted Arrays - Brute Force Solution")

	fmt.Printf("Example 1: %v, m: %d, %v, n: %d\nOutput: %v\n", case1_nums1, case1_m, case1_nums2, case1_n, bruteForceMergeSortedArrays(case1_nums1, case1_m, case1_nums2, case1_n))
	fmt.Printf("Example 2: %v, m: %d, %v, n: %d\nOutput: %v\n", case2_nums1, case2_m, case2_nums2, case2_n, bruteForceMergeSortedArrays(case2_nums1, case2_m, case2_nums2, case2_n))
	fmt.Printf("Example 3: %v, m: %d, %v, n: %d\nOutput: %v\n", case3_nums1, case3_m, case3_nums2, case3_n, bruteForceMergeSortedArrays(case3_nums1, case3_m, case3_nums2, case3_n))

	// reassign variables to test two pointer solution
	case1_nums1 = []int{1, 2, 3, 0, 0, 0}
	case1_nums2 = []int{2, 5, 6}

	case2_nums1 = []int{1}
	case2_nums2 = []int{}

	case3_nums1 = []int{0}
	case3_nums2 = []int{1}

	println("\nMerge Sorted Arrays - Two Pointers Solution")

	fmt.Printf("Example 1: %v, m: %d, %v, n: %d\nOutput: %v\n", case1_nums1, case1_m, case1_nums2, case1_n, twoPointerMergeSortedArrays(case1_nums1, case1_m, case1_nums2, case1_n))
	fmt.Printf("Example 2: %v, m: %d, %v, n: %d\nOutput: %v\n", case2_nums1, case2_m, case2_nums2, case2_n, twoPointerMergeSortedArrays(case2_nums1, case2_m, case2_nums2, case2_n))
	fmt.Printf("Example 3: %v, m: %d, %v, n: %d\nOutput: %v\n", case3_nums1, case3_m, case3_nums2, case3_n, twoPointerMergeSortedArrays(case3_nums1, case3_m, case3_nums2, case3_n))
}
