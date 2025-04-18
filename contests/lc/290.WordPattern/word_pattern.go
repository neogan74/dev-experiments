package main

import "strings"

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(words) != len(pattern) {
		return false
	}

	charToWord := make(map[byte]string)
	wordToChar := make(map[string]byte)

	for i := 0; i < len(pattern); i++ {
		c := pattern[i]
		w := words[i]

		if mappedWord, exists := charToWord[c]; exists {
			if mappedWord != w {
				return false
			}
		} else {
			charToWord[c] = w
		}

		if mappedChar, exists := wordToChar[w]; exists {
			if mappedChar != c {
				return false
			}
		} else {
			wordToChar[w] = c
		}
	}
	return true
}
