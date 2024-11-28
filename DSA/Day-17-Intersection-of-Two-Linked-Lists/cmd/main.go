package main

import "fmt"

/*
Given the heads of two singly linked-lists headA and headB, return the node at which the two lists intersect.
If the two linked lists have no intersection at all, return null.

The goal is to find the node at which two singly linked lists intersect.
If the two linked lists have no intersection, return nil.

Given the heads of two singly linked lists, headA and headB, return the node at which they intersect, or nil if they do not intersect.

Example:
Input:
List A: 4 -> 1 -> 8 -> 4 -> 5 -> nil
List B: 5 -> 6 -> 1 -> 8 -> 4 -> 5 -> nil

Output:
Intersecting Node: Node with value 8.

Key Insights:
	- The two linked lists may differ in length before their intersection point.
	- After the intersection, they share all remaining nodes.
	- If we traverse both lists entirely and switch heads when reaching the end of one list, they will align at the intersection point or reach the end (nil) together.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nIntersection of Two Linked Lists - Leetcode #160")

	twoPointerSolution()

	hashmapSolution()
}
