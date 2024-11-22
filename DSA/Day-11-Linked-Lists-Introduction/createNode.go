package main

// Node represents a single element in the linked list
type Node struct {
	data int   // value stored in the node
	next *Node // pointer to the next node
}

// append a node to the end of the linked list
func appendNode(head *Node, data int) *Node {
	// create a new node via the memory address
	newNode := &Node{data: data}

	// if the linked list is empty, the new node becomes the head
	if head == nil {
		return newNode
	}

	// traverse the linked list to find the last node
	current := head
	for current.next != nil {
		current = current.next
	}

	// append the new node
	current.next = newNode

	return head
}
