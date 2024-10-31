package main

import "fmt"

/*
A singly linked list has nodes that contain a value and a pointer to the next node.
*/

// Node struct for singly linked list
type Node struct {
	Value int
	Next  *Node
}

// LinkedList struct to manage the list
type LinkedList struct {
	Head *Node
}

// add to the end of the linked list
func (l *LinkedList) Append(value int) {
	newNode := &Node{Value: value}
	if l.Head == nil {
		l.Head = newNode
	} else {
		current := l.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
}

// display linked list
func (l *LinkedList) Display() {
	current := l.Head
	for current != nil {
		fmt.Printf("%d -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}

// implement LinkedList algorithm
func SinglLinkedList() {
	fmt.Println("\nSingly Linked List")
	list := LinkedList{}

	list.Append(1)
	list.Append(2)
	list.Append(3)

	list.Display()
}

/*
A doubly linked list has nodes that contain pointers to both the next and previous nodes.
*/
type DoubleNode struct {
	Value    int
	Next     *DoubleNode
	Previous *DoubleNode
}

type DoublyLinkedList struct {
	Head *DoubleNode
}

func (l *DoublyLinkedList) Append(value int) {
	newNode := &DoubleNode{Value: value}
	if l.Head == nil {
		l.Head = newNode
	} else {
		current := l.Head
		for current.Next != nil {
			current = current.Next

		}
		current.Next = newNode
		newNode.Previous = current
	}
}

// Display prints the elements in the list in the "1 -> 2 -> 3 -> nil" format
func (l *DoublyLinkedList) Display() {
	current := l.Head
	for current != nil {
		fmt.Printf("%d -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}

func DoubleLinkedList() {
	fmt.Println("\nDouble Linked List")

	dlist := DoublyLinkedList{}

	dlist.Append(1)
	dlist.Append(2)
	dlist.Append(3)
	dlist.Display()
}

/*
The Display method starts at the Head node and iterates over each node in the list.
For each node, it prints the node's Value followed by " -> ".
Once it reaches the end (where current is nil), it prints "nil" to indicate the end of the list.
*/
