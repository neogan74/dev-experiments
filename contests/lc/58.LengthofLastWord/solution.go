package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	// Обрезаем пробелы с начала и конца строки.
	s = strings.TrimSpace(s)
	length := 0

	// Проходим по строке с конца до первого пробела.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			break
		}
		length++
	}
	return length
}

func main() {
	// Примеры:
	s1 := "Hello World"
	fmt.Printf("Длина последнего слова в %q: %d\n", s1, lengthOfLastWord(s1)) // Ожидается 5

	s2 := "   fly me   to   the moon  "
	fmt.Printf("Длина последнего слова в %q: %d\n", s2, lengthOfLastWord(s2)) // Ожидается 4

	s3 := "luffy is still joyboy"
	fmt.Printf("Длина последнего слова в %q: %d\n", s3, lengthOfLastWord(s3)) // Ожидается 6
}
