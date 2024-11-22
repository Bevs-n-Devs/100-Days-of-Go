package main

import "fmt"

// travers the linked list and print each nodes data
func traverseList(head *Node) {
	currentNode := head
	for currentNode != nil {
		fmt.Print(currentNode.data, " -> ")
		currentNode = currentNode.next
	}
	fmt.Println("nil")
}
