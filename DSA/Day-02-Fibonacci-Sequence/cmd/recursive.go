package main

func fibonacciRecursive(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
}
