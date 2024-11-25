package main

import "fmt"

/*
1. Create a dummy node to act as the start of the merged list.
2. Use a pointer current to build the merged list.
3. Compare the current nodes of both lists:
	- Append the smaller node to the merged list.
	- Advance the pointer of the list from which the node was taken.
4. If one list is exhausted, append the remaining nodes of the other list.
5. Return the next node of the dummy node as the head of the merged list.
*/

// append the node to the end of the linked list
func appendNode(head *ListNode, value int) *ListNode {
	newNode := &ListNode{Val: value}

	if head == nil {
		return newNode
	}

	current := head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode

	return head
}

// merge two sorted linked lists iteratively
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{} // dummy node to act as the start of the merged list
	current := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			// add list1's current node to the merged list
			current.Next = list1
			list1 = list1.Next
		} else {
			// add list2's current node to the merged list
			current.Next = list2
			list2 = list2.Next
		}

		current = current.Next
	}

	// append the memaining nodes of the non-empty list
	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	// return the merged list (skipping  the dummy node)
	return dummy.Next
}

// traverse the linked list and print each nodes data
func traverseList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}

func iterativeSolution() {
	// iterative solution
	fmt.Println("\nIterative Solution")

	// create two soreted linked lists
	var list1, list2 *ListNode
	values1 := []int{1, 2, 4}
	values2 := []int{1, 3, 4}

	for _, value := range values1 {
		list1 = appendNode(list1, value)
	}

	for _, value := range values2 {
		list2 = appendNode(list2, value)
	}

	// print the orginal lists
	fmt.Println("Original Linked List 1:")
	traverseList(list1)
	fmt.Println("Original Linked List 2:")
	traverseList(list2)

	// merge the two lists
	mergedList := mergeTwoLists(list1, list2)

	// print the merged list
	fmt.Println("\nMerged Linked List:")
	traverseList(mergedList)
}
