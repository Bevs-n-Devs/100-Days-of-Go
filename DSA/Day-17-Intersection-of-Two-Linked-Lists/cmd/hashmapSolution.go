package main

import (
	"fmt"
)

/*
1. Store Nodes in a Hash Map:
	- Traverse the first list and store each node in a hash map.
2. Check for Intersection:
	- Traverse the second list and check if any node exists in the hash map.
3. Return the First Match:
	- If a match is found, return that node as the intersection.
*/

func findIntersectionHashMap(headA, headB *ListNode) *ListNode {
	visited := make(map[*ListNode]bool)

	// traverse the 1st list and store the nodes in a hash map
	for current := headA; current != nil; current = current.Next {
		visited[current] = true
	}

	// traverse the 2nd list and check if any node exists in the hash map
	for current := headB; current != nil; current = current.Next {
		if visited[current] {
			// if a match is found, return that node as the intersection
			return current
		}
	}

	return nil
}

func hashmapSolution() {
	fmt.Println("\nHashmap Solution")

	// create intersecting node
	intersectionNode := &ListNode{Val: 8, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}

	// create list A: 4 -> 1 -> 8 -> 4 -> 5
	headA := appendNode(nil, 4)
	headA = appendNode(headA, 1)
	lastA := headA
	for lastA.Next != nil {
		lastA = lastA.Next
	}
	lastA.Next = intersectionNode

	// create list B: 5 -> 6 -> 1 -> 8 -> 4 -> 5
	headB := appendNode(nil, 5)
	headB = appendNode(headB, 6)
	headB = appendNode(headB, 1)
	lastB := headB
	for lastB.Next != nil {
		lastB = lastB.Next
	}
	lastB.Next = intersectionNode

	// print both lists
	fmt.Println("List A:")
	traverseList(headA)
	fmt.Println("List B:")
	traverseList(headB)

	// find intersection using hash map
	intersection := findIntersectionHashMap(headA, headB)
	if intersection != nil {
		fmt.Printf("Intersection: Node with value %d\n", intersection.Val)
	} else {
		fmt.Println("No intersection found.")
	}
}
