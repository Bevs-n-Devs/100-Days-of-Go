package main

import "fmt"

/*
We can use a hash map to store visited nodes. If a node is encountered again during traversal, a cycle exists.

Steps:
	1. Traverse the linked list.
	2. Use a hash map (or map in Go) to store references to visited nodes.
	3. If a node is visited more than once, return true.
	4. If the traversal reaches nil, return false.
*/

func hashmapSolution(head *ListNode) bool {
	// hahmap solution
	fmt.Println("\nHashmap Solution")

	// create a map to store visited nodes
	visited := make(map[*ListNode]bool)
	current := head
	for current != nil {
		if visited[current] {
			// node is visted, so cycle exists
			return true
		}

		// mark the current node as visited
		visited[current] = true
		current = current.Next
	}

	// if the traversal reaches nil, return false
	return false
}
