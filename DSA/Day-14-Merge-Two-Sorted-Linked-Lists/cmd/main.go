package main

import "fmt"

/*
You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.

The goal is to merge two sorted linked lists into a single sorted linked list while preserving the order.

Example:

Input:
List1: 1 -> 3 -> 5 -> 7 -> nil
List2: 2 -> 4 -> 6 -> 8 -> nil

Output:
Merged List: 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> nil
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nDay 14 of DSA: Merge Two Sorted Linked Lists - Leetcode #21")

	iterativeSolution()

	recursiveSolution()
}
