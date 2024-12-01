package main

import "fmt"

func recursiveAddTwoNumbers(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}

	sum := carry
	if l1 != nil {
		sum += l1.Val
		l1 = l1.Next
	}
	if l2 != nil {
		sum += l2.Val
		l2 = l2.Next
	}

	node := &ListNode{Val: sum % 10}
	node.Next = recursiveAddTwoNumbers(l1, l2, sum/10)
	return node
}

func recursiveSolution() {
	fmt.Println("\nRecursive Solution")

	fmt.Println("\nRecursive Solution")

	// Example input: 342 + 465
	l1 := createLinkedList([]int{2, 4, 3}) // Represents 342
	l2 := createLinkedList([]int{5, 6, 4}) // Represents 465

	fmt.Println("Adding Two Numbers:")
	fmt.Print("List 1: ")
	printLinkedList(l1)
	fmt.Print("List 2: ")
	printLinkedList(l2)

	result := recursiveAddTwoNumbers(l1, l2, 0)
	fmt.Print("Result: ")
	printLinkedList(result)

}
