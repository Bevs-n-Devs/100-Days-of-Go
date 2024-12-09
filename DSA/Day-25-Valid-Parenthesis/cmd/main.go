package main

import "fmt"

/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

1. Open brackets must be closed by the same type of brackets.
2. Open brackets must be closed in the correct order.
3. Every close bracket has a corresponding open bracket of the same type.
*/

func isValid(s string) bool {
	stack := []rune{}
	matching := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		switch char {
		case '(', '{', '[':
			// push opening brackets onto the stack
			stack = append(stack, char)
		case ')', '}', ']':
			// check if the stack is empty
			if len(stack) == 0 {
				return false
			}

			// pop the top of the stack and compare
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if top != matching[char] {
				return false
			}
		}
	}
	// if the stack is empty, all brackest are matched
	return len(stack) == 0
}

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nDay 25 of DSA: Valid Parenthesis - Leetcode #20")

	fmt.Println(isValid("()"))     // Output: true
	fmt.Println(isValid("()[]{}")) // Output: true
	fmt.Println(isValid("(]"))     // Output: false
	fmt.Println(isValid("([)]"))   // Output: false
	fmt.Println(isValid("{[]}"))   // Output: true
}
