package main

import (
	"fmt"
)

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}

	// Инициализируем границы обхода.
	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1

	result := []int{}

	for top <= bottom && left <= right {
		// 1. Верхняя строка: слева направо.
		for j := left; j <= right; j++ {
			result = append(result, matrix[top][j])
		}
		top++ // верхнюю границу сдвигаем вниз

		// 2. Правый столбец: сверху вниз.
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right-- // правую границу сдвигаем влево

		// 3. Нижняя строка: справа налево (если еще осталась строка)
		if top <= bottom {
			for j := right; j >= left; j-- {
				result = append(result, matrix[bottom][j])
			}
			bottom-- // сдвигаем нижнюю границу вверх
		}

		// 4. Левый столбец: снизу вверх (если еще остались столбцы)
		if left <= right {
			for i := bottom; i >= top; i-- {
				result = append(result, matrix[i][left])
			}
			left++ // сдвигаем левую границу вправо
		}
	}

	return result
}

func main() {
	// Пример: матрица 3х3
	matrix1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Printf("Spiral order: %v\n", spiralOrder(matrix1))

	// Пример: матрица 3х4
	matrix2 := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Printf("Spiral order: %v\n", spiralOrder(matrix2))
}
