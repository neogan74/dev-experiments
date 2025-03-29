package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	var cleaned strings.Builder
	for _, char := range strings.ToLower(s) {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			cleaned.WriteRune(char)
		}
	}
	cleanedStr := cleaned.String()
	for i := 0; i < len(cleanedStr)/2; i++ {
		if cleanedStr[i] != cleanedStr[len(cleanedStr)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	s := "A man, a plan, a canal: Panama"
	res := isPalindrome(s)
	fmt.Println(res, " (expected true)") //true
	s2 := "race a car"
	res2 := isPalindrome(s2)
	fmt.Println(res2, " (expected false)") //false
	s3 := " "
	res3 := isPalindrome(s3)
	fmt.Println(res3, " (expected true)") //true
}
