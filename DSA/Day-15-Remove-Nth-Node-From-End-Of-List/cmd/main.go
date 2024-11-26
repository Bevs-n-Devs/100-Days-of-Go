package main

import "fmt"

/*
Given the head of a linked list, remove the nth node from the end of the list and return its head.

Given the head of a linked list and an integer n, remove the N-th node from the end of the list and return the head of the modified list.

Example:
Input:
Linked List: 1 -> 2 -> 3 -> 4 -> 5 -> nil
n = 2

Output:
Modified Linked List: 1 -> 2 -> 3 -> 5 -> nil

Key Insight:
To remove the N-th node from the end in one pass:
	1. Use two pointers: fast and slow.
	2. Advance fast by n + 1 steps ahead of slow.
	3. Move both pointers one step at a time until fast reaches the end of the list.
	4. slow will now point to the node before the target node. Adjust the next pointer of slow to remove the target node.

Steps:
1. Create a Dummy Node:
	- A dummy node is useful to handle edge cases, like removing the head node, without extra logic.
2. Advance Fast Pointer:
	- Move the fast pointer n + 1 steps ahead of slow.
3. Traverse Until End:
	- Move fast and slow one step at a time until fast reaches the end.
4. Adjust Pointers:
	- Set slow.next to slow.next.next to remove the N-th node.
*/

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

// remove the Nth node from the end of the list
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	fast, slow := dummy, dummy

	// move fast pointer n+1 steps ahead of slow pointer
	for i := 0; i <= n; i++ {
		if fast == nil {
			return head // if n is larger than the length of the list
		}
		fast = fast.Next
	}

	// move both fast and slow pointers until fast reaches the end
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// remove the Nth node from the end
	slow.Next = slow.Next.Next

	return dummy.Next // return new head
}

func traverseList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Print(current.Val, " -> ")
		current = current.Next
	}
	fmt.Println("nil")
}

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nDay 15 of DSA: Remove Nth Node From End of List - Leetcode #19")

	// create a linked list
	var head *ListNode
	values := []int{1, 2, 3, 4, 5}

	// add nodes to the end of the linked list
	for _, value := range values {
		head = appendNode(head, value)
	}

	// print the orginal list
	fmt.Println("Original Linked List:")
	traverseList(head)

	// remove the 2nd node from the end
	n := 2
	head = removeNthFromEnd(head, n)

	// print the modified list
	fmt.Printf("\nList after removing the %dth node from the end: \n", n)
	traverseList(head)

}
