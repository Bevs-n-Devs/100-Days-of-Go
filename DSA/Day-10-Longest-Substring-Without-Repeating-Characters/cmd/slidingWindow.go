package main

/*
The sliding window approach uses two pointers to represent a window of characters. It moves the window to ensure all characters within it are unique.

Steps:
1. Use a hash map to store the last seen index of each character.
2. Use two pointers (start and end) to represent the window.
3. If a character is repeated, move the start pointer to exclude the earlier occurrence.
4. Update the maximum length after processing each character.
*/

func slidingWindow(s string) int {
	// sliding window solution
	charIndex := make(map[byte]int) // map to store last seen index of each character
	maxLength := 0

	start := 0 // start of the current window

	for end := 0; end < len(s); end++ {
		if lastIndex, found := charIndex[s[end]]; found && lastIndex >= start {
			// if character already exists in the window, move the start pointer
			start = lastIndex + 1
		}

		// update the last index of the current character
		charIndex[s[end]] = end

		// calculate the max length of the current window
		maxLength = max(maxLength, end-start+1) // using max helper function from brute force solution
	}

	return maxLength
}
