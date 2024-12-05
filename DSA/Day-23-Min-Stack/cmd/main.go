package main

import "fmt"

type MinStack struct {
	stack    []int
	minStack []int
}

// Constructor initializes the Min Stack
func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}

// Push adds an element to the stack
func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	// Push the minimum value onto minStack
	if len(this.minStack) == 0 || val <= this.minStack[len(this.minStack)-1] {
		this.minStack = append(this.minStack, val)
	}
}

// Pop removes the top element from the stack
func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	top := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	// If the popped value is the minimum, pop from minStack as well
	if len(this.minStack) > 0 && top == this.minStack[len(this.minStack)-1] {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
}

// Top retrieves the top element of the stack
func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 0 // or some error value
	}
	return this.stack[len(this.stack)-1]
}

// GetMin retrieves the minimum element in the stack
func (this *MinStack) GetMin() int {
	if len(this.minStack) == 0 {
		return 0 // or some error value
	}
	return this.minStack[len(this.minStack)-1]
}

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nDay 23 of DSA: Min Stack - Leetcode #155")

	stack := Constructor()

	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
	fmt.Println(stack.GetMin()) // Output: -3

	stack.Pop()
	fmt.Println(stack.Top())    // Output: 0
	fmt.Println(stack.GetMin()) // Output: -2
}
