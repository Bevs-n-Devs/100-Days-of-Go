package main

import "fmt"

// node represents a single element in the linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// append a node to the end of the linked list
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

// reverse a linked likst and return the new head
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	currentNode := head

	for currentNode != nil {
		nextTemp := currentNode.Next
		currentNode.Next = prev
		prev = currentNode
		currentNode = nextTemp
	}
	return prev
}

// checks if a linked list is a palindrome
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true // empty or single-node lists are palindromes
	}

	// STEP 1: find the middle of the linked list
	// create pointers
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// STEP 2: reverse the second half of the linked list
	secondHalf := reverseList(slow)

	// STEP 3: compare the first half with the reversed second half
	firstHalf, secondHalfCopy := head, secondHalf

	for secondHalfCopy != nil {
		if firstHalf.Val != secondHalfCopy.Val {
			return false
		}
		firstHalf = firstHalf.Next
		secondHalfCopy = secondHalfCopy.Next
	}

	// STEP 4: restore the linked list
	reverseList(secondHalf)

	return true
}

// traverse the linked list and print each nodes data
func traverseList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nDay 16 of DSA: Palindrome Linked List - Leetcode #234")

	// create a linked list
	var head *ListNode
	values := []int{1, 2, 3, 2, 1}

	// add nodes to the end of the linked list
	for _, value := range values {
		head = appendNode(head, value)
	}

	// print the orginal list
	fmt.Println("Original Linked List:")
	traverseList(head)

	// check if the linked list is a palindrome
	fmt.Println("\nIs the linked list a palindrome?", isPalindrome(head))

	// create a non-palindrome linked list
	var head2 *ListNode
	values2 := []int{1, 2, 3, 4, 5}

	// add nodes to the end of the linked list
	for _, value := range values2 {
		head2 = appendNode(head2, value)
	}

	// print the orginal list
	fmt.Println("\nOriginal Linked List:")
	traverseList(head2)

	// check if the linked list is a palindrome
	fmt.Println("\nIs the linked list a palindrome?", isPalindrome(head2))
}
