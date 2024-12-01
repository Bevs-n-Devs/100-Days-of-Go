package main

import "fmt"

/*
A linked list is given where each node contains an additional pointer, random,
which may point to any node in the list (or null). You need to create a deep copy of this list, such that:
	1. Each node in the new list contains the same value as the corresponding node in the original list.
	2. Each random pointer in the new list points to the copied node corresponding to the node that the random pointer in the original list pointed to.

Example:
Input: head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
Output: [[7,null],[13,0],[11,4],[10,2],[1,0]]
*/

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// print linked list with random pointers
func printList(head *Node) {
	current := head
	for current != nil {
		randomVal := "nil"
		if current.Random != nil {
			randomVal = fmt.Sprintf("%d", current.Random.Val)
		}
		fmt.Printf("Node(%d, random: %s) -> ", current.Val, randomVal)
		current = current.Next
	}
	fmt.Println("nil")
}

// create a linked list with random pointers from a given structure
func createList(values []int, randomIndices []int) *Node {
	nodes := make([]*Node, len(values))
	for i, val := range values {
		nodes[i] = &Node{Val: val}
		if i > 0 {
			nodes[i-1].Next = nodes[i]
		}
	}

	for i, randomIdx := range randomIndices {
		if randomIdx != -1 {
			nodes[i].Random = nodes[randomIdx]
		}
	}
	return nodes[0]
}

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nDay 20 of DSA: Copy List With Random Pointer - Leetcode #138")

	bruteForceHashmapSolution()

	optimalSolution()
}
