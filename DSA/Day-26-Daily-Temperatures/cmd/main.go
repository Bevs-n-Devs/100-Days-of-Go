package main

import "fmt"

/*
Given an array of integers temperatures represents the daily temperatures,
return an array answer such that answer[i] is the number of days you have
to wait after the ith day to get a warmer temperature. If there is no future
day for which this is possible, keep answer[i] == 0 instead.

Example:

Input: temperatures = [73,74,75,71,69,72,76,73]
Output: [1,1,4,2,1,1,0,0]

*/

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nDaily Temperatures - Leetcode #739")

	bruteFoceSolution()

	stackSolution()
}
