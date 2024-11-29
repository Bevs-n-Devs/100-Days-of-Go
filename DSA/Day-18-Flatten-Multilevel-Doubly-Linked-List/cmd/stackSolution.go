package main

import "fmt"

func flattenStack(head *Node) *Node {
	if head == nil {
		return nil
	}

	stack := make([]*Node, 0) // stack to store Next nodes durign traversal
	current := head

	for current != nil {
		// if th enode has a Child, we need to flatten it
		if current.Child != nil {
			// if theres a next node, push it onto the stack
			if current.Next != nil {
				stack = append(stack, current.Next)
			}

			// connect current node to the child list
			current.Next = current.Child
			current.Child.Prev = current
			current.Child = nil // remove the child pointer
		}

		// if we reach the end of sublist and stack not empty, pop the next node
		if current.Next == nil && len(stack) > 0 {
			next := stack[len(stack)-1]  // get the last element of the stack
			stack = stack[:len(stack)-1] // remove the last element of the stack
			current.Next = next          // connect the current node to the next node
			next.Prev = current          // connect the next node to the current node
		}

		// move to the next node
		current = current.Next
	}

	return head

}

func stackSolution() {
	fmt.Println("\nStack Solution")

	// Create nodes for the first level
	head := createNode(1)
	node2 := createNode(2)
	node3 := createNode(3)
	node4 := createNode(4)
	node5 := createNode(5)
	node6 := createNode(6)

	// Link the first level
	head.Next = node2
	node2.Prev = head
	node2.Next = node3
	node3.Prev = node2
	node3.Next = node4
	node4.Prev = node3
	node4.Next = node5
	node5.Prev = node4
	node5.Next = node6
	node6.Prev = node5

	// Create nodes for the second level
	node7 := createNode(7)
	node8 := createNode(8)
	node9 := createNode(9)
	node10 := createNode(10)
	node3.Child = node7
	node7.Next = node8
	node8.Prev = node7
	node8.Next = node9
	node9.Prev = node8
	node9.Next = node10
	node10.Prev = node9

	// Create nodes for the third level
	node11 := createNode(11)
	node12 := createNode(12)
	node8.Child = node11
	node11.Next = node12
	node12.Prev = node11

	// Print the original multilevel list
	fmt.Println("Original Multilevel Doubly Linked List:")
	printList(head)

	// Flatten the list iteratively
	flattened := flattenStack(head)

	// Print the flattened list
	fmt.Println("Flattened Doubly Linked List:")
	printList(flattened)
}
