package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for _, s := range strs[1:] {
		i := 0
		// Сравниваем текущий префикс и новую строку посимвольно
		for ; i < len(prefix) && i < len(s); i++ {
			if prefix[i] != s[i] {
				break
			}
		}
		// Уменьшаем префикс до совпавшей части
		prefix = prefix[:i]

		// Если префикс стал пустым, останавливаемся
		if prefix == "" {
			return ""
		}
	}

	return prefix
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))          // "fl"
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))             // ""
	fmt.Println(longestCommonPrefix([]string{"interview", "internet", "internal"})) // "inte"
}
