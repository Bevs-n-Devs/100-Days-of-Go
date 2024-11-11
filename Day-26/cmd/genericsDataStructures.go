package main

import "fmt"

/*
Generics are very useful in data structures where flexibility is essential, such as for collections like slices, stacks, or queues.

Stack[T any] defines a stack for any type T.
Push and Pop work for any type, allowing us to create stacks for integers, strings, and more.
*/

// Stack struct with a generic type parameter T
type Stack[T any] struct {
	elements []T
}

// push method to add an element
func (s *Stack[T]) Push(value T) {
	s.elements = append(s.elements, value)
}

// pop method to remove and return the last element
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false // return zero and false if stack is empty
	}

	val := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return val, true
}

func genericsWithDataStructures() {
	fmt.Println("\nGenerics With Data Structures")

	instStack := Stack[int]{}
	instStack.Push(10)
	instStack.Push(20)
	fmt.Println(instStack.Pop()) // 20, true
	fmt.Println(instStack.Pop()) // 10, true

	stringStack := Stack[string]{}
	stringStack.Push("Hello")
	stringStack.Push("world")
	fmt.Println(stringStack.Pop()) // world, true
}
