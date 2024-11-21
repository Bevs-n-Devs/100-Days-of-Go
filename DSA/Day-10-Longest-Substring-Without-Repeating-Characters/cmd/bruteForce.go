package main

/*
The brute force solution examines all possible substrings, checking each for uniqueness.

Steps:
1. Generate all substrings of s using two nested loops.
2. Check if each substring contains all unique characters using a helper function.
3. Keep track of the maximum length encountered.
*/

// helper function to check if a substring contains all unique characters
func isUnique(s string, start, end int) bool {
	// create a set to store unique characters
	charSet := make(map[byte]bool)

	// iterate through the substring and add each character to the set
	for i := start; i <= end; i++ {
		// check if the character is already in the set
		if charSet[s[i]] {
			return false
		}
		charSet[s[i]] = true
	}
	return true
}

// helper function to find the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func bruteForce(s string) int {
	// brute force solution
	maxLength := 0

	// iterate through all possible substrings
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if isUnique(s, i, j) {
				maxLength = max(maxLength, j-i+1)
			}
		}
	}

	return maxLength
}
