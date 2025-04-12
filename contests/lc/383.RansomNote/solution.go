package _83_RansomNote

import "fmt"

func one() string {
	return fmt.Sprintln("cool")
}

func canConstruct(ransomNote string, magazine string) bool {
	// Создаем счётчик для букв из magazine
	counts := make([]int, 26)
	for _, ch := range magazine {
		counts[ch-'a']++
	}
	// Проверяем, достаточно ли каждой буквы для ransomNote
	for _, ch := range ransomNote {
		counts[ch-'a']--
		if counts[ch-'a'] < 0 {
			return false
		}
	}
	return true
}
