package main

import "fmt"

type MyQueue struct {
	inStack  []int
	outStack []int
}

// Constructor initializes the queue
func Constructor() MyQueue {
	return MyQueue{
		inStack:  []int{},
		outStack: []int{},
	}
}

// Push adds an element to the queue
func (this *MyQueue) Push(x int) {
	this.inStack = append(this.inStack, x)
}

// Pop removes and returns the front element of the queue
func (this *MyQueue) Pop() int {
	if len(this.outStack) == 0 {
		// Transfer all elements from inStack to outStack
		for len(this.inStack) > 0 {
			top := this.inStack[len(this.inStack)-1]
			this.inStack = this.inStack[:len(this.inStack)-1]
			this.outStack = append(this.outStack, top)
		}
	}
	// Pop the top element from outStack
	top := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	return top
}

// Peek returns the front element of the queue without removing it
func (this *MyQueue) Peek() int {
	if len(this.outStack) == 0 {
		// Transfer all elements from inStack to outStack
		for len(this.inStack) > 0 {
			top := this.inStack[len(this.inStack)-1]
			this.inStack = this.inStack[:len(this.inStack)-1]
			this.outStack = append(this.outStack, top)
		}
	}
	return this.outStack[len(this.outStack)-1]
}

// Empty checks if the queue is empty
func (this *MyQueue) Empty() bool {
	return len(this.inStack) == 0 && len(this.outStack) == 0
}

func main() {
	fmt.Println("Hello world, hello Yaw")

	fmt.Println("Day 22 of DSA: Implement Queue - Leetcode #232")

	queue := Constructor()

	queue.Push(1)
	queue.Push(2)
	fmt.Println(queue.Peek())  // Output: 1
	fmt.Println(queue.Pop())   // Output: 1
	fmt.Println(queue.Empty()) // Output: false
}
