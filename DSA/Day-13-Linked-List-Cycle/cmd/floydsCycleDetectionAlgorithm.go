package main

import "fmt"

/*
This is a more efficient approach that uses two pointers:

1. A slow pointer that moves one step at a time.
2. A fast pointer that moves two steps at a time.

Key Insight:
	- If there’s a cycle, the fast pointer will eventually meet the slow pointer.
	- If there’s no cycle, the fast pointer will reach nil.

Steps:
	1. Initialize two pointers (slow and fast) at the head.
	2. Move slow by one step and fast by two steps in each iteration.
	3. If slow and fast meet, a cycle exists; return true.
	4. If fast or fast.next becomes nil, there’s no cycle; return false.
*/

func floydsCycleDetectionAlgorithm(head *ListNode) bool {
	// floyds cycle detection algorithm
	fmt.Println("\nFloyd's Cycle Detection Algorithm")

	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head

	// loop will keep going as long as there is a next node in the list.
	for fast != nil && fast.Next != nil {
		slow = slow.Next      // move slow by one step
		fast = fast.Next.Next // move fast by two steps

		if slow == fast {
			// cycle detected
			return true
		}
	}

	// no cycle detected
	return false
}
