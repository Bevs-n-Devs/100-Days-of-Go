package main

import "fmt"

func bruteForceHashmap(head *Node) *Node {
	if head == nil {
		return nil
	}

	// step 1 : create a map to store the copied nodes (ignoring random pointers)
	oldToNew := make(map[*Node]*Node)
	current := head
	for current != nil {
		oldToNew[current] = &Node{Val: current.Val}
		current = current.Next
	}

	// step 2 : update th enext & random pointers of the copied nodes
	current = head
	for current != nil {
		if current.Next != nil {
			oldToNew[current].Next = oldToNew[current.Next]
		}
		if current.Random != nil {
			oldToNew[current].Random = oldToNew[current.Random]
		}
		current = current.Next
	}

	// return th edeep copied head
	return oldToNew[head]
}

func bruteForceHashmapSolution() {
	fmt.Println("\nBrute Force Hashmap Solution")

	// values and random indices
	values := []int{7, 13, 11, 10, 1}
	randomIndices := []int{-1, 0, 4, 2, 0}

	head := createList(values, randomIndices)
	fmt.Println("Original Linked List:")
	printList(head)

	// deep copy using brute force hashmap solution
	copiedHead := bruteForceHashmap(head)
	fmt.Println("\nCopied Linked List:")
	printList(copiedHead)
}
