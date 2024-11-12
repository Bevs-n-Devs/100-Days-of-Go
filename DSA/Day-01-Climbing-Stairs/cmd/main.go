package main

import "fmt"

/*
You are climbing a staircase. It takes 𝑛 steps to reach the top.
Each time you can either climb 1 or 2 steps.
In how many distinct ways can you climb to the top?

Example:

Input: n = 3
Output: 3

Explanation: There are three ways to climb to the top:
			 1 step + 1 step + 1 step
			 1 step + 2 steps
			 2 steps + 1 step
*/

// returns the number of ways to climb n stairs
func climbStairs(n int) int {
	// check numbers of stairs is equal to 1
	if n == 1 {
		return 1
	}

	// check if num stairs equal to 2
	if n == 2 {
		return 2
	}

	// initial first three steps
	oneStepBefore := 1  // ways to reach step 1
	twoStepsBefore := 2 // ways to rech step 2
	result := oneStepBefore + twoStepsBefore

	// loop from steps 2 up to n, using previous 2 steps
	for i := 2; i < n; i++ {
		result = oneStepBefore + twoStepsBefore

		// move twoStepsBefore by one step
		twoStepsBefore = oneStepBefore

		// update the oneStepBefore to be the result
		oneStepBefore = result
	}

	// result will be the number of ways to reach n stairs
	return result
}

func main() {
	fmt.Println("Hello world, hello Yaw!")

	fmt.Println("\nClimbing Stairs Challenge - Leetcode #70")

	case1 := 2
	fmt.Printf("Number of ways to climb %v stairs: %v\n", case1, climbStairs(case1))

	case2 := 3
	fmt.Printf("Number of ways to climb %v stairs: %v\n", case2, climbStairs(case2))

	case3 := 4
	fmt.Printf("Number of ways to climb %d stairs: %d\n", case3, climbStairs(case3))

	case4 := 5
	fmt.Printf("Number of ways to climb %d stairs: %d\n", case4, climbStairs(case4))
}
