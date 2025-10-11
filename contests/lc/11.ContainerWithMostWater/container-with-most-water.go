package main

import (
	"fmt"
)

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0

	for left < right {
		// Вычисляем текущую площадь контейнера.
		var currentArea int
		if height[left] < height[right] {
			currentArea = height[left] * (right - left)
		} else {
			currentArea = height[right] * (right - left)
		}

		// Обновляем максимум, если нашли больше.
		if currentArea > maxArea {
			maxArea = currentArea
		}

		// Сдвигаем указатель с меньшей высотой.
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

func main() {
	// Пример:
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Printf("Максимальная площадь: %d\n", maxArea(height)) // Ожидаемый результат: 49
}
