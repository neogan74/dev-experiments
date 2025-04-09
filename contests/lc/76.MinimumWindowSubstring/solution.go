package main

import (
	"fmt"
	"math"
)

// minWindow возвращает минимальную подстроку строки s, содержащую все символы строки t.
func minWindow(s string, t string) string {
	if len(t) == 0 || len(s) < len(t) {
		return ""
	}

	// Шаг 1. Подсчитываем символы строки t.
	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	// Количество уникальных символов, которых требуется достичь.
	required := len(need)

	// currentCounts хранит частоты символов текущего окна.
	currentCounts := make(map[byte]int)
	formed := 0 // количество символов, удовлетворяющих требуемую частоту

	l, r := 0, 0
	// Для ответа: (длина окна, левый индекс, правый индекс)
	ansLen, ansL, ansR := math.MaxInt32, 0, 0

	for r < len(s) {
		// Добавляем символ s[r] в текущее окно.
		c := s[r]
		currentCounts[c]++
		// Если текущая частота символа совпадает с требуемой, увеличиваем formed.
		if count, ok := need[c]; ok && currentCounts[c] == count {
			formed++
		}

		// Когда окно валидно (содержит все символы с нужными частотами)
		for l <= r && formed == required {
			// Обновляем ответ, если текущее окно меньше предыдущего.
			if r-l+1 < ansLen {
				ansLen = r - l + 1
				ansL = l
				ansR = r
			}

			// Пытаемся уменьшить окно, исключая s[l].
			c = s[l]
			currentCounts[c]--
			// Если после удаления s[l] частота стала меньше требуемой, окно перестаёт быть валидным.
			if count, ok := need[c]; ok && currentCounts[c] < count {
				formed--
			}
			l++
		}

		r++
	}

	if ansLen == math.MaxInt32 {
		return ""
	}
	return s[ansL : ansR+1]
}

func main() {
	// Пример 1:
	s1 := "ADOBECODEBANC"
	t1 := "ABC"
	fmt.Printf("s = %q, t = %q → %q\n", s1, t1, minWindow(s1, t1)) // Ожидается "BANC"

	// Пример 2:
	s2 := "a"
	t2 := "a"
	fmt.Printf("s = %q, t = %q → %q\n", s2, t2, minWindow(s2, t2)) // Ожидается "a"

	// Пример 3:
	s3 := "a"
	t3 := "aa"
	fmt.Printf("s = %q, t = %q → %q\n", s3, t3, minWindow(s3, t3)) // Ожидается ""
}
