package main

import (
	"fmt"
	"strconv"
)

/*
You are given an array of strings tokens that represents an arithmetic expression in a Reverse Polish Notation.

Evaluate the expression. Return an integer that represents the value of the expression.

Note that:

The valid operators are '+', '-', '*', and '/'.
Each operand may be an integer or another expression.
The division between two integers always truncates toward zero.
There will not be any division by zero.
The input represents a valid arithmetic expression in a reverse polish notation.
The answer and all the intermediate calculations can be represented in a 32-bit integer.


Example:

Input: tokens = ["2","1","+","3","*"]
Output: 9
Explanation: ((2 + 1) * 3) = 9
*/

func evalRPN(tokens []string) int {
	stack := []int{}

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			// Pop the top two numbers from the stack
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// Perform the operation and push the result
			var result int
			switch token {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				result = a / b
			}
			stack = append(stack, result)
		default:
			// Token is a number, convert to int and push onto the stack
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}

	// The final result is the only number left in the stack
	return stack[0]
}

func main() {
	fmt.Println("\nHello world, hello Yaw")

	fmt.Println("\nDay 24 of DSA: Evaluate Reverse Polish Notation - Leetcode #150")

	tokens := []string{"2", "1", "+", "3", "*"}
	result := evalRPN(tokens)
	fmt.Println(result) // Output: 9

	tokens2 := []string{"4", "13", "5", "/", "+"}
	result2 := evalRPN(tokens2)
	fmt.Println(result2) // Output: 6
}
