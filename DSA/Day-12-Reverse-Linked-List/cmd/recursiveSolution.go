package main

import "fmt"

/*
In the recursive approach:
	1. Break the problem into smaller parts: reverse the rest of the list and then fix the current node's link.
	2. Use the base case: if the list is empty or has only one node, return the head.
	3. Reverse the list recursively by updating the next pointer of the next node.
	4. The last node in the original list becomes the new head.
*/

func reverseListRecursive(head *ListNode) *ListNode {
	// recursive solution

	// base case: if the list is empty or has one node, return the head
	if head == nil || head.Next == nil {
		return head
	}

	// reverse the rest of the list
	newHead := reverseListRecursive(head.Next)

	// fux the current node's link
	head.Next.Next = head
	head.Next = nil

	return newHead
}

func recursiveSolution() {
	fmt.Println("\nRecursive Solution")

	// Input values to create the linked list
	values := []int{10, 20, 30, 40, 50}

	// Create the linked list
	var head *ListNode
	for _, value := range values {
		head = appendNode(head, value)
	}

	// Print the original list
	fmt.Println("Original Linked List:")
	traverseList(head)

	// Reverse the linked list recursively
	head = reverseListRecursive(head)

	// Print the reversed list
	fmt.Println("Reversed Linked List (Recursive):")
	traverseList(head)

}
