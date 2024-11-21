package main

import "fmt"

/*
Given a string s, find the length of the longest substring without repeating characters.

Example:
Input: s = "abcabcbb"
Output: 3
*/

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nDay 10 of DSA: Longest Substring Without Repeating Characters - Leetcode #3")

	case1 := "abcabcbb"
	case2 := "bbbbb"
	case3 := "pwwkew"

	fmt.Println("\nSolution 1: Brute Force Solution")
	fmt.Printf("Example 1: %s\nOutput: %d\n", case1, bruteForce(case1))
	fmt.Printf("Example 2: %s\nOutput: %d\n", case2, bruteForce(case2))
	fmt.Printf("Example 3: %s\nOutput: %d\n", case3, bruteForce(case3))

	fmt.Println("\nSolution 2: Sliding Window Solution")
	fmt.Printf("Example 1: %s\nOutput: %d\n", case1, slidingWindow(case1))
	fmt.Printf("Example 2: %s\nOutput: %d\n", case2, slidingWindow(case2))
	fmt.Printf("Example 3: %s\nOutput: %d\n", case3, slidingWindow(case3))

}
