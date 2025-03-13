package main

import "fmt"

// numberOfSubstrings returns the number of substrings that contain at least one occurrence of 'a', 'b', and 'c'.
func numberOfSubstrings(s string) int {
	n := len(s)
	// Frequency array for 'a', 'b', and 'c'.
	// index 0 for 'a', 1 for 'b', 2 for 'c'
	count := [3]int{}
	result := 0
	left := 0

	// Iterate over the string with the right pointer
	for right := 0; right < n; right++ {
		// Update the count for the current character.
		count[s[right]-'a']++

		// Check if the current window [left, right] contains at least one of each character.
		for count[0] > 0 && count[1] > 0 && count[2] > 0 {
			// All substrings from the current window ending at any index from right to n-1 are valid.
			result += n - right

			// Shrink the window from the left and update the count.
			count[s[left]-'a']--
			left++
		}
	}

	return result
}

func main() {
	// Test cases
	fmt.Println(numberOfSubstrings("abcabc")) // Expected output: 10
	fmt.Println(numberOfSubstrings("aaacb"))  // Expected output: 3
	fmt.Println(numberOfSubstrings("abc"))    // Expected output: 1
}
