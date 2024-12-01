package main

import "fmt"

func digitByDigitAddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	current := dummy
	carry := 0 // carry for the next digit when the sum is greater than 9

	// traverse both linked lists
	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry

		// add the values from both lists (if available)
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		// update carry for next iteration
		carry = sum / 10

		// create a new node with the current digit
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	return dummy.Next

}

func digitByDigitSolution() {
	fmt.Println("\nDigit by Digit Solution")

	// Example input: 342 + 465
	l1 := createLinkedList([]int{2, 4, 3}) // Represents 342
	l2 := createLinkedList([]int{5, 6, 4}) // Represents 465

	fmt.Println("Adding Two Numbers:")
	fmt.Print("List 1: ")
	printLinkedList(l1)
	fmt.Print("List 2: ")
	printLinkedList(l2)

	result := digitByDigitAddTwoNumbers(l1, l2)
	fmt.Print("Result: ")
	printLinkedList(result)
}
