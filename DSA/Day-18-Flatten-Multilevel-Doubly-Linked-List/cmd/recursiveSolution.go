package main

import "fmt"

// flatten recursively
func flattenRecursive(head *Node) *Node {
	if head == nil {
		return nil
	}

	// create a dummy head to manage the previous pointer
	dummy := &Node{Val: 0, Next: head}

	// helper function to flatten the linked list recursively - depth first search
	var dfs func(node *Node) *Node
	dfs = func(node *Node) *Node {
		current := node
		var last *Node

		for current != nil {
			next := current.Next
			// if the node has a child, flatten it
			if current.Child != nil {
				childLast := dfs(current.Child) // flatten the child list

				// connect the current node to the child
				current.Next = current.Child
				current.Child.Prev = current
				current.Child = nil // remove the child pointer

				// connect the last node of the flttened child to the original next
				if next != nil {
					childLast.Next = next
					next.Prev = childLast
				}

				last = childLast // update the last to point to the end of the child list

			} else {
				last = current // update last to the current node
			}
			current = next // move to the next node
		}

		return last
	}

	// flatten the list starting from the head
	dfs(head)

	// disconnect/remove the dummy head
	dummy.Next.Prev = nil
	return dummy.Next
}

// traverse and print doubly linked list
func printList(head *Node) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" <-> ")
		}
		head = head.Next
	}
	fmt.Println()
}

// helper function to create a new node
func createNode(val int) *Node {
	return &Node{Val: val}
}

func recursiveSolution() {
	fmt.Println("\nRecursive Solution")

	// create nodes for the 1st level
	head := createNode(1)
	node2 := createNode(2)
	node3 := createNode(3)
	node4 := createNode(4)
	node5 := createNode(5)
	node6 := createNode(6)

	// link thw 1st level
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

	// create nodes for the 2nd level
	node7 := createNode(7)
	node8 := createNode(8)
	node9 := createNode(9)
	node10 := createNode(10)
	node3.Child = node7
	node7.Prev = node3
	node7.Next = node8
	node8.Prev = node7
	node8.Next = node9
	node9.Prev = node8
	node9.Next = node10
	node10.Prev = node9

	// create nodes for the 3rd level
	node11 := createNode(11)
	node12 := createNode(12)
	node8.Child = node11
	node11.Next = node12
	node12.Prev = node11

	// print the original multilevel list
	fmt.Println("Original Multilevel Doubly Linked List:")
	printList(head)

	// flatten the multilevel list
	flattened := flattenRecursive(head)

	// print the flattened list
	fmt.Println("Flattened Multilevel Doubly Linked List:")
	printList(flattened)

}
