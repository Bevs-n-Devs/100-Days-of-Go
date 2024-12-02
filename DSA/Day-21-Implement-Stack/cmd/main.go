package main

import "fmt"

type MyStack struct {
	queue []int
}

// Constructor initializes the stack
func Constructor() MyStack {
	return MyStack{queue: []int{}}
}

// Push adds an element onto the stack
func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x) // Add to the end of the queue
	n := len(this.queue)
	// Rotate the queue so the new element appears at the front
	for i := 0; i < n-1; i++ {
		this.queue = append(this.queue, this.queue[0])
		this.queue = this.queue[1:]
	}
}

// Pop removes and returns the top element
func (this *MyStack) Pop() int {
	if this.Empty() {
		return -1 // Error value; stack is empty
	}
	top := this.queue[0]
	this.queue = this.queue[1:] // Remove the first element
	return top
}

// Top returns the top element without removing it
func (this *MyStack) Top() int {
	if this.Empty() {
		return -1 // Error value; stack is empty
	}
	return this.queue[0]
}

// Empty checks if the stack is empty
func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nDay 21 of DSA: Implement Stack - Leetcode #225")

	// Initialize the stack
	obj := Constructor()

	// Test Push operation
	obj.Push(1)
	fmt.Println("Pushed 1")
	obj.Push(2)
	fmt.Println("Pushed 2")

	// Test Top operation
	fmt.Println("Top element:", obj.Top()) // Output: 2

	// Test Pop operation
	fmt.Println("Popped element:", obj.Pop()) // Output: 2

	// Test Empty operation
	fmt.Println("Is stack empty?", obj.Empty()) // Output: false

	// Pop remaining element
	fmt.Println("Popped element:", obj.Pop()) // Output: 1

	// Check if stack is empty again
	fmt.Println("Is stack empty?", obj.Empty()) // Output: true
}
