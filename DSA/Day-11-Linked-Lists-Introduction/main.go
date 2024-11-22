package main

import "fmt"

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nLinked Lists Introduction")

	// input values to create the linked list
	values := []int{1, 2, 3, 4, 5}

	// create the linked list
	var head *Node
	for _, value := range values {
		head = appendNode(head, value)
	}

	// traverse the linked list
	fmt.Println("\nTraversing the linked list:")
	traverseList(head)
}
