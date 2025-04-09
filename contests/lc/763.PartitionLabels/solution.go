package main

import (
	"fmt"
)

func partitionLabels(s string) []int {
	// Шаг 1. Находим последнее вхождение каждого символа.
	lastIndex := make(map[rune]int)
	for i, ch := range s {
		lastIndex[ch] = i
	}

	var partitions []int
	start, end := 0, 0

	// Шаг 2. Проходим по строке.
	for i, ch := range s {
		// Обновляем конец участка.
		if lastIndex[ch] > end {
			end = lastIndex[ch]
		}
		// Если достигли конца текущего участка, фиксируем его.
		if i == end {
			partitions = append(partitions, end-start+1)
			start = i + 1
		}
	}

	return partitions
}

func main() {
	// Пример из условия:
	s := "ababcbacadefegdehijhklij"
	result := partitionLabels(s)
	fmt.Println("Разбиение строки:", result) // Ожидаемый вывод: [9 7 8]
}
