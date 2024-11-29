package main

import "fmt"

/*
You are given a doubly linked list, which contains nodes that have a next pointer, a previous pointer,
and an additional child pointer. This child pointer may or may not point to a separate doubly linked list,
also containing these special nodes. These child lists may have one or more children of their own, and so
on, to produce a multilevel data structure as shown in the example below.

Given the head of the first level of the list, flatten the list so that all the nodes appear in a single-level,
doubly linked list. Let curr be a node with a child list. The nodes in the child list should appear after curr
and before curr.next in the flattened list.

Return the head of the flattened list. The nodes in the list must have all of their child pointers set to null.

In this problem, we aim to "flatten" a multilevel doubly linked list into a single-level doubly linked list.
Each node in the list may have a child pointer pointing to another doubly linked list. These lists can be nested multiple levels deep.


Definition:
Each node in the doubly linked list has the following attributes:
	- Val: 		An integer value.
	- Prev: 	A pointer to the previous node.
	- Next: 	A pointer to the next node.
	- Child: 	A pointer to the head of a child doubly linked list (can be nil).


Task:
Flatten the multilevel doubly linked list into a single-level doubly linked list, where:
	- All nodes appear in a depth-first traversal order.
	- The child pointer is set to nil for all nodes.

Example:
	Input: head = [1,2,3,4,5,6,null,null,null,7,8,9,10,null,null,11,12]
	Output: [1,2,3,7,8,11,12,9,10,4,5,6]

Input: head = [1,2,3,4,5,6,7,8,9,10,11,12]
1 -- 2 -- 3 -- 4 -- 5 -- 6
          |
          7 -- 8 -- 9 -- 10
                |
                11 -- 12

Output: [1,2,3,7,8,11,12,9,10,4,5,6]
1 -- 2 -- 3 -- 7 -- 8 -- 11 -- 12 -- 9 -- 10 -- 4 -- 5 -- 6
*/

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nDay 18 of DSA: Flatten Multilevel Doubly Linked List - Leetcode #430")

	recursiveSolution()

	stackSolution()
}
