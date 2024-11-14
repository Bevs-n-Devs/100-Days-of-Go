package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nDay 2 of DSA: Fibonacci Sequence - Leetcode #509")

	var case1, case2, case3, case4 int
	case1 = 2
	case2 = 3
	case3 = 4
	case4 = 5

	fmt.Println("\nFibonacci Sequence: Recursive Solution")
	fmt.Printf("Fibonacci (recursive) of %d: %d\n", case1, fibonacciRecursive(case1))
	fmt.Printf("Fibonacci (recursive) of %d: %d\n", case2, fibonacciRecursive(case2))
	fmt.Printf("Fibonacci (recursive) of %d: %d\n", case3, fibonacciRecursive(case3))
	fmt.Printf("Fibonacci (recursive) of %d: %d\n", case4, fibonacciRecursive(case4))

	fmt.Println("\nFibonacci Sequence: Top-Down (Memoization) Solution")
	fmt.Printf("Fibonacci (memoization) of %d: %d\n", case1, fibonacciMemoized(case1))
	fmt.Printf("Fibonacci (memoization) of %d: %d\n", case2, fibonacciMemoized(case2))
	fmt.Printf("Fibonacci (memoization) of %d: %d\n", case3, fibonacciMemoized(case3))
	fmt.Printf("Fibonacci (memoization) of %d: %d\n", case4, fibonacciMemoized(case4))

	fmt.Println("\nFibonacci Sequence: Down-Up (Iterative) Solution")
	fmt.Printf("Fibonacci (iterative) of %d: %d:\n", case1, fibonacciIterative(case1))
	fmt.Printf("Fibonacci (iterative) of %d: %d:\n", case2, fibonacciIterative(case2))
	fmt.Printf("Fibonacci (iterative) of %d: %d:\n", case3, fibonacciIterative(case3))
	fmt.Printf("Fibonacci (iterative) of %d: %d:\n", case4, fibonacciIterative(case4))

}
