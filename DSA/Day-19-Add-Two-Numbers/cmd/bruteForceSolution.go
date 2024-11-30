package main

import "fmt"

func bruteForceAddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// convert linked list to integers
	num1 := 0
	multipler := 1
	for l1 != nil {
		num1 += l1.Val * multipler // add the value of the current node to the result
		multipler *= 10            // multiply by 10 for each digit
		l1 = l1.Next               // move to the next node
	}

	num2 := 0
	multipler = 1
	for l2 != nil {
		num2 += l2.Val * multipler // add the value of the current node to the result
		multipler *= 10            // multiply by 10 for each digit
		l2 = l2.Next               // move to the next node
	}

	// add the two numbers
	sum := num1 + num2

	// convert the sum back to a linked list
	dummy := &ListNode{Val: 0}
	current := dummy

	for sum > 0 {
		current.Next = &ListNode{Val: sum % 10} // add the current digit to the linked list
		sum /= 10                               // divide by 10 to get the next digit
		current = current.Next                  // move to the next node
	}

	return dummy.Next
}

func bruteForceSolution() {
	// brute force solution
	fmt.Println("\nBrute Force Solution")

	// Example input: 342 + 465
	l1 := createLinkedList([]int{2, 4, 3}) // Represents 342
	l2 := createLinkedList([]int{5, 6, 4}) // Represents 465

	fmt.Println("Adding Two Numbers:")
	fmt.Print("List 1: ")
	printLinkedList(l1)
	fmt.Print("List 2: ")
	printLinkedList(l2)

	result := bruteForceAddTwoNumbers(l1, l2)
	fmt.Print("Result: ")
	printLinkedList(result)
}
