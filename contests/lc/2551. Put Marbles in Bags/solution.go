package main

import (
	"fmt"
	"sort"
)

func putMarbles(marbles []int, k int) int64 {
	n := len(marbles)
	// Если k равно 1, значит разбиения нет, а стоимость фиксирована – разность 0.
	if k == 1 {
		return 0
	}

	// Вычисляем массив diffs длины n-1:
	// diffs[i] = marbles[i] + marbles[i+1]
	diffs := make([]int, n-1)
	for i := 1; i < n; i++ {
		diffs[i-1] = marbles[i-1] + marbles[i]
	}

	// Сортируем diffs по возрастанию
	sort.Ints(diffs)

	// Нам нужно выбрать m = k-1 границ.
	m := k - 1
	var minExtra, maxExtra int64 = 0, 0
	// Сумма минимальных m значений
	for i := 0; i < m; i++ {
		minExtra += int64(diffs[i])
	}
	// Сумма максимальных m значений
	for i := len(diffs) - m; i < len(diffs); i++ {
		maxExtra += int64(diffs[i])
	}

	return maxExtra - minExtra
}

func main() {
	// Пример 1:
	// Пусть marbles = [1, 3, 5, 1] и k = 2.
	// diffs: [1+3=4, 3+5=8, 5+1=6] → отсортированные: [4,6,8]
	// m = 1, maxExtra = 8, minExtra = 4, ответ = 8-4 = 4.
	marbles1 := []int{1, 3, 5, 1}
	k1 := 2
	fmt.Printf("Marbles: %v, k = %d → Разность: %d\n", marbles1, k1, putMarbles(marbles1, k1))

	// Пример 2:
	// Пусть marbles = [3, 2, 1, 4, 5] и k = 2.
	// diffs: [3+2=5, 2+1=3, 1+4=5, 4+5=9] → отсортированные: [3,5,5,9]
	// m = 1, maxExtra = 9, minExtra = 3, ответ = 9-3 = 6.
	marbles2 := []int{3, 2, 1, 4, 5}
	k2 := 2
	fmt.Printf("Marbles: %v, k = %d → Разность: %d\n", marbles2, k2, putMarbles(marbles2, k2))
}
