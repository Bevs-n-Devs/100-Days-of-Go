package main

import (
	"errors"
	"fmt"
)

// define the stack structure
type Stack struct {
	elements []int
}

// push method adds an element to the top of the stack
func (s *Stack) Push(value int) {
	s.elements = append(s.elements, value)
}

// pop method removes and returns the top of the stack
func (s *Stack) Pop() (int, error) {
	if len(s.elements) == 0 {
		return 0, fmt.Errorf("Stack is empty")
	}
	val := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return val, nil
}

func ImplementStack() {
	fmt.Println("\nImplementing A Stack")

	// craete an empty stack
	stack := Stack{}

	// push elements onto the stack
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	// print stack content after pushes
	fmt.Println("Stack after pushes:", stack.elements)

	// pop element from the stack
	for i := 0; i < 3; i++ {
		val, err := stack.Pop()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}
		fmt.Println("Popped value:", val)
	}

	// try to pop from an empty stack to demonstate handling
	_, err := stack.Pop()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

type Queue struct {
	elements []rune
}

// enqueue add an element to the end of the queue
func (q *Queue) Enqueue(value rune) {
	q.elements = append(q.elements, value)
}

// dequeue removes an element from the front of the queue
func (q *Queue) Dequeue() (rune, error) {
	if len(q.elements) == 0 {
		return 0, errors.New("Queue is empty")
	}
	val := q.elements[0]
	q.elements = q.elements[1:]
	return val, nil
}

// isBalanced function using a stack (for comparison)
func isBalanced(s string) bool {
	stack := []rune{}
	matching := map[rune]rune{')': '(', '}': '{', ']': '['}

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else if char == ')' || char == '}' || char == ']' {
			if len(stack) == 0 || stack[len(stack)-1] != matching[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func ImplementQueue() {
	fmt.Println("\nImplement Queue")

	// create empty queue
	queue := Queue{}

	queue.Enqueue('a')
	queue.Enqueue('b')
	queue.Enqueue('c')

	fmt.Println("Queue after enqueus:", string(queue.elements)) // output: [a b c]

	for i := 0; i < 3; i++ {
		val, err := queue.Dequeue()
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Dequeeud:", string(val))
		}
	}

	// test isBalanced with a stack (stack implementation)
	fmt.Println("Is balanced:", isBalanced("{[()]}"))
}
