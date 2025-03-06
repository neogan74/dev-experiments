package main

import (
	"fmt"
)

// mergeStrings merges two strings in alternating order.
func mergeAlternately(word1, word2 string) string {
	result := ""
	i := 0
	// Find the minimum length between the two strings.
	minLen := len(word1)
	if len(word2) < minLen {
		minLen = len(word2)
	}
	// Append characters from both strings alternately.
	for i < minLen {
		result += string(word1[i])
		result += string(word2[i])
		i++
	}

	// Append remaining letters from the longer string.
	if i < len(word1) {
		result += word1[i:]
	}
	if i < len(word2) {
		result += word2[i:]
	}

	return result
}

func main() {
	fmt.Println(mergeAlternately("abc", "pqr")) // Output: "apbqcr"
	fmt.Println(mergeAlternately("ab", "pqrs")) // Output: "apbqrs"
	fmt.Println(mergeAlternately("abcd", "pq")) // Output: "apbqcd"
}
