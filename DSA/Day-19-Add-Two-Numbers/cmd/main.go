package main

import "fmt"

/*
You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order, and each of their nodes contains a single digit.
Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// create linked likst helper function
func createLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	current := head

	for _, num := range nums[1:] {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}

	return head
}

// print linked list helper function
func printLinkedList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next // shift the head to the next node
	}
	fmt.Println()

}

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nDay 19 of DSA: Add Two Numbers - Leetcode #2")

	bruteForceSolution()

	digitByDigitSolution()

	recursiveSolution()
}
