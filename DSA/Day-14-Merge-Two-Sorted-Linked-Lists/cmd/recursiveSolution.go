package main

import "fmt"

// merges two sorted linked lists recursively
func mergeTwoListsRecursive(list1 *ListNode, list2 *ListNode) *ListNode {
	// base cases: if one list is empty, return the other
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	// compare the current nodes of both lists
	if list1.Val < list2.Val {
		list1.Next = mergeTwoListsRecursive(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoListsRecursive(list1, list2.Next)
		return list2
	}
}

func recursiveSolution() {
	fmt.Println("\nRecursive Solution")

	// create two sorted linked lists
	var list1, list2 *ListNode
	values1 := []int{1, 2, 4}
	values2 := []int{1, 3, 4}

	for _, value := range values1 {
		list1 = appendNode(list1, value)
	}

	for _, value := range values2 {
		list2 = appendNode(list2, value)
	}

	// print the original lists
	fmt.Println("Original Linked List 1:")
	traverseList(list1)
	fmt.Println("Original Linked List 2:")
	traverseList(list2)

	// merge the two lists recursively
	mergedList := mergeTwoListsRecursive(list1, list2)

	// print the merged list
	fmt.Println("\nMerged Linked List:")
	traverseList(mergedList)
}
