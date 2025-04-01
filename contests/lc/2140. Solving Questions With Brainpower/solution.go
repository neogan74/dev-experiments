package main

import (
	"fmt"
)

func mostPoints(questions [][]int) int64 {
	n := len(questions)
	// dp[i] - максимальное количество баллов, которое можно набрать, начиная с вопроса i.
	dp := make([]int64, n+1) // dp[n] = 0 по умолчанию

	// Идем с конца к началу
	for i := n - 1; i >= 0; i-- {
		points := int64(questions[i][0])
		brainpower := questions[i][1]
		nextIndex := i + brainpower + 1
		scoreIfSolved := points
		if nextIndex < n {
			scoreIfSolved += dp[nextIndex]
		}
		// Выбираем максимальный вариант: решить вопрос или пропустить его.
		if scoreIfSolved > dp[i+1] {
			dp[i] = scoreIfSolved
		} else {
			dp[i] = dp[i+1]
		}
	}
	return dp[0]
}

func main() {
	// Пример 1:
	questions1 := [][]int{
		{3, 2},
		{4, 3},
		{4, 4},
		{2, 5},
	}
	// Ожидаемый результат: 5 (пример из условия, если оно дано)
	fmt.Println("Максимальное количество баллов (пример 1):", mostPoints(questions1))

	// Пример 2:
	questions2 := [][]int{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 4},
		{5, 5},
	}
	fmt.Println("Максимальное количество баллов (пример 2):", mostPoints(questions2))
}
