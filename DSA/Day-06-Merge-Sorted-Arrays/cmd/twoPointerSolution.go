package main

/*
Pointer Initialization:
- p1 starts at the last meaningful element of nums1 (index m - 1).
- p2 starts at the last element of nums2 (index n - 1).
- p starts at the end of nums1 (index m + n - 1).

Merging in Reverse:
- Compare nums1[p1] and nums2[p2]. Place the larger value at nums1[p] and decrement the respective pointer.
- If p1 or p2 runs out of elements, stop comparing.

Copy Remaining Elements:
- If any elements remain in nums2, copy them to nums1.
- We don't need to copy elements from nums1, as they are already in place.
*/

func twoPointerMergeSortedArrays(nums1 []int, m int, nums2 []int, n int) []int {
	// set up pointers
	p1 := m - 1    // pointer for the last element of nums1
	p2 := n - 1    // pointer for the last element of nums2
	p := m + n - 1 // pointer for the end of nums1

	// merge in reverse order
	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1-- // move pointer for nums1 to the previous element
		} else {
			nums1[p] = nums2[p2]
			p2-- // move pointer for nums2 to the previous element
		}
		p-- // move pointer for nums1 to the previous element
	}

	// copy remaining elements from nums2 to nums1
	for p2 >= 0 {
		nums1[p] = nums2[p2]
		p2-- // move pointer for nums2 to the previous element
		p--  // move pointer for nums1 to the previous element
	}

	return nums1
}
