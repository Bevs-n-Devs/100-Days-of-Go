package main

import "fmt"

func copyRandomListOptimal(head *Node) *Node {
	if head == nil {
		return nil
	}

	// step 1: Interweave the original and copied nodes
	current := head
	for current != nil {
		copy := &Node{Val: current.Val}
		copy.Next = current.Next
		current.Next = copy
		current = copy.Next
	}

	// step 2: assign random pointer for the copied nodes
	current = head
	for current != nil {
		if current.Random != nil {
			current.Next.Random = current.Random.Next
		}
		current = current.Next.Next
	}

	// step 3: separate into original and copied lists
	current = head
	dummy := &Node{Val: 0}
	copyCurrent := dummy
	for current != nil {
		copy := current.Next
		copyCurrent.Next = copy
		copyCurrent = copy // move to the next copied node
		current.Next = copy.Next
		current = copy.Next // move to the next original node
	}

	return dummy.Next
}

func optimalSolution() {
	fmt.Println("\nOptimal Solution")

	// values and random indices
	values := []int{7, 13, 11, 10, 1}
	randomIndices := []int{-1, 0, 4, 2, 0}

	head := createList(values, randomIndices)
	fmt.Println("Original Linked List:")
	printList(head)

	// deep copy using brute force hashmap solution
	copiedHead := copyRandomListOptimal(head)
	fmt.Println("\nCopied Linked List:")
	printList(copiedHead)
}
