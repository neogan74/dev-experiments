package main

import (
	"fmt"
)

// isSubsequence проверяет, является ли s подпоследовательностью t.
func isSubsequence(s string, t string) bool {
	i, j := 0, 0
	// Проходим по строке t, пытаясь найти символы строки s в нужном порядке.
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	// Если все символы s найдены (i == len(s)), значит s является подпоследовательностью t.
	return i == len(s)
}

func main() {
	// Пример 1:
	s1 := "abc"
	t1 := "ahbgdc"
	fmt.Printf("s = %q, t = %q → %v\n", s1, t1, isSubsequence(s1, t1)) // Ожидается: true

	// Пример 2:
	s2 := "axc"
	t2 := "ahbgdc"
	fmt.Printf("s = %q, t = %q → %v\n", s2, t2, isSubsequence(s2, t2)) // Ожидается: false
}
