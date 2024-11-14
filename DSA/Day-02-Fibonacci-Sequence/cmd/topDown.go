package main

// store the calculated Fibonacci numbers
var memo = map[int]int{}

func fibonacciMemoized(n int) int {
	if n <= 1 {
		return n
	}

	// check if value already exists in memo
	// When you index a map in Go you get two return values; the second one (which is optional) is a boolean that indicates if the key exists.
	if val, exists := memo[n]; exists {
		return val
	}

	// calculate and store result in memo
	memo[n] = fibonacciMemoized(n-1) + fibonacciMemoized(n-2)
	return memo[n]
}
