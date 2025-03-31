package main

import (
	"fmt"
)

// twoSum находит два числа в отсортированном массиве numbers,
// сумма которых равна target, и возвращает их индексы (1-based).
func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			// возвращаем индексы с 1-based нумерацией
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++ // если сумма меньше, двигаем левый указатель вправо
		} else {
			right-- // если сумма больше, двигаем правый указатель влево
		}
	}
	return nil // по условию задачи гарантируется, что решение существует
}

func main() {
	// Пример:
	numbers := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(numbers, target)
	fmt.Printf("Найденные индексы: %v\n", result) // Ожидается: [1, 2]
}
