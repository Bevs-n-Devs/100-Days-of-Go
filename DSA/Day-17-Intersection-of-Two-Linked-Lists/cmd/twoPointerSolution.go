package main

import "fmt"

/*
1. Two Pointers:
	- Use two pointers, pA and pB, starting at the heads of headA and headB.
2. Traverse the Lists:
	- When a pointer reaches the end of its list, redirect it to the head of the other list.
	- If there is an intersection, the two pointers will eventually meet at the intersecting node.
	- If there is no intersection, the two pointers will both become nil.
3. Convergence:
	- The two pointers align after at most m + n steps, where m and n are the lengths of the two lists.
*/

// add node to the end of the linked list
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

// find intersection using two pointers
// this returns the node at which the teo linked lists intersect, nil if no intersection is found
func findIntersectionTwoPointers(headA, headB *ListNode) *ListNode {
	// return nil if any of the linked lists is empty
	if headA == nil || headB == nil {
		return nil
	}

	// initialize two pointers to the heads of the linked lists
	pA, pB := headA, headB

	// traverse both lists until they point to the same node
	for pA != pB {
		// move to the next node in the list, or switch to the other list's head
		if pA == nil {
			pA = headB // switch to the other list's head if pA reaches the end
		} else {
			pA = pA.Next // move to the next node in the first list
		}

		if pB == nil {
			pB = headA // switch to the other list's head if pB reaches the end
		} else {
			pB = pB.Next // move to the next node in the second list
		}
	}

	// pA and pB will either meet at the intersection or at nil (at the end of one of the lists)
	return pA
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

func twoPointerSolution() {
	fmt.Println("\nTwo Pointer Solution")

	// create the two linked lists with an intersection
	intersectionNode := &ListNode{Val: 8, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}

	// list A: 4 -> 1 -> 8 -> 4 -> 5
	headA := appendNode(nil, 4)
	headA = appendNode(headA, 1)
	lastA := headA
	for lastA.Next != nil {
		lastA = lastA.Next
	}
	lastA.Next = intersectionNode

	// list B: 5 -> 6 -> 1 -> 8 -> 4 -> 5
	headB := appendNode(nil, 5)
	headB = appendNode(headB, 6)
	headB = appendNode(headB, 1)
	lastB := headB
	for lastB.Next != nil {
		lastB = lastB.Next
	}
	lastB.Next = intersectionNode

	// print both lists
	fmt.Println("List A:")
	traverseList(headA)
	fmt.Println("List B:")
	traverseList(headB)

	// find and print the intersection
	intersection := findIntersectionTwoPointers(headA, headB)
	if intersection != nil {
		fmt.Printf("Intersection Node: %d\n", intersection.Val)
	} else {
		fmt.Println("No intersection found")
	}
}
