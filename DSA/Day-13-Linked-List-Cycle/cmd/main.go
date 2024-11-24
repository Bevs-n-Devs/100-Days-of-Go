package main

import "fmt"

/*
Given head, the head of a linked list, determine if the linked list has a cycle in it.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer.
Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.

Return true if there is a cycle in the linked list. Otherwise, return false.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// Append a node to the end of the linked list
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

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nDay 13 of DSA: Linked List Cycle - Leetcode #141")

	case1 := []int{3, 2, 0, -4}
	case2 := []int{1, 2}

	// Create a linked list with a cycle
	var head *ListNode
	values := case1
	for _, value := range values {
		head = appendNode(head, value)
	}

	// Introduce a cycle: Point the last node to the second node
	if head != nil {
		current := head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = head.Next // Create cycle
	}

	// Check for cycle
	fmt.Println("Cycle Detected:", hashmapSolution(head))
	fmt.Println("Cycle Detected:", floydsCycleDetectionAlgorithm(head))

	// Create a linked list without a cycle
	var head2 *ListNode
	values2 := case2
	for _, value := range values2 {
		head2 = appendNode(head2, value)
	}

	// Check for cycle
	fmt.Println("Cycle Detected:", hashmapSolution(head2))
	fmt.Println("Cycle Detected:", floydsCycleDetectionAlgorithm(head2))
}
