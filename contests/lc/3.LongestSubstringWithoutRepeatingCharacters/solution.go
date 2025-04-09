package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	// Мапа для хранения последнего индекса, где встретился символ.
	lastIndex := make(map[rune]int)
	maxLength := 0
	left := 0

	// Преобразуем строку в срез рун для корректной работы с юникодом.
	runes := []rune(s)

	for right, ch := range runes {
		// Если символ уже встречался и его последний индекс >= left,
		// перемещаем left за этот индекс.
		if prev, ok := lastIndex[ch]; ok && prev >= left {
			left = prev + 1
		}
		// Обновляем последний индекс текущего символа.
		lastIndex[ch] = right
		// Вычисляем длину текущего окна.
		if currentLength := right - left + 1; currentLength > maxLength {
			maxLength = currentLength
		}
	}
	return maxLength
}

func main() {
	// Примеры:
	s1 := "abcabcbb" // Ожидаемая длина: 3 ("abc")
	s2 := "bbbbb"    // Ожидаемая длина: 1 ("b")
	s3 := "pwwkew"   // Ожидаемая длина: 3 ("wke")
	fmt.Printf("s = %q, длина = %d\n", s1, lengthOfLongestSubstring(s1))
	fmt.Printf("s = %q, длина = %d\n", s2, lengthOfLongestSubstring(s2))
	fmt.Printf("s = %q, длина = %d\n", s3, lengthOfLongestSubstring(s3))
}
