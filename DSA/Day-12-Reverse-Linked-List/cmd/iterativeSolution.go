package main

import "fmt"

/*
1. Start with three pointers:
	- prev initialized to nil (to track the reversed portion).
	- current initialized to the head (to track the node being processed).
	- next to hold the next node temporarily.
2. Update the next pointer of the current node to point to prev (reversing the link).
3. Move prev and current one step forward.
4. Repeat until current becomes nil.
5. The new head of the reversed list is prev.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// append a node to the ned of the linked list
func appendNode(head *ListNode, value int) *ListNode {
	newNode := &ListNode{Val: value}

	if head == nil {
		return newNode
	}

	currentNode := head
	// loop will keep going as long as there is a next node in the list.
	for currentNode.Next != nil {
		currentNode = currentNode.Next // move to the next node
	}

	currentNode.Next = newNode
	return head
}

// reverse the linked list iteratively
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	currentNode := head

	for currentNode != nil {
		// save th enext node
		next := currentNode.Next

		// reverse the link
		currentNode.Next = prev

		// move pointers one step forward
		prev = currentNode
		currentNode = next
	}

	// new head of the reversed list
	return prev
}

// traverse the linked list and print each nodes data
func traverseList(head *ListNode) {
	currentHead := head
	for currentHead != nil {
		fmt.Printf("%d -> ", currentHead.Val)
		currentHead = currentHead.Next
	}

	fmt.Println("nil")
}

func iterativeSolution() {
	// iterative solution
	fmt.Println("\nIterative Solution")

	// input values to create the linked list
	values := []int{10, 20, 30, 40, 50}

	// create the linked list
	var head *ListNode
	for _, value := range values {
		head = appendNode(head, value)
	}

	// print the orginal list
	fmt.Println("Original Linked List:")
	traverseList(head)

	// reverse the linked list
	head = reverseList(head)

	// print the reversed list
	fmt.Println("\nReversed Linked List:")
	traverseList(head)
}
