package main

import (
	"fmt"
)

func rotate(matrix [][]int) {
	n := len(matrix)

	// Шаг 1: транспонируем матрицу
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// Шаг 2: переворачиваем каждую строку
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("Before rotation:")
	printMatrix(matrix)

	rotate(matrix)

	fmt.Println("\nAfter rotation:")
	printMatrix(matrix)
}
