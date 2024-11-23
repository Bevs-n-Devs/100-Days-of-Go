package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func dynamicDataStructures() {
	fmt.Println("\nDynamic Data Structures")

	info :=
		`- Pointers are essential for building linked lists, trees, and other dynamic data structures.
`
	fmt.Println(info)

	n1 := &Node{Value: 1}
	n2 := &Node{Value: 2}
	n3 := &Node{Value: 3}
	n1.Next = n2 // linking the nodes
	n2.Next = n3

	fmt.Println("Node 1 Value:", n1.Value)
	fmt.Println("Node 2 Value:", n1.Next.Value)
	fmt.Println("Node 3 Value:", n1.Next.Next.Value)

	fmt.Println("\nTraversing the linked list:")
	fmt.Printf("%d -> %d -> %d\n", n1.Value, n1.Next.Value, n1.Next.Next.Value)
}
