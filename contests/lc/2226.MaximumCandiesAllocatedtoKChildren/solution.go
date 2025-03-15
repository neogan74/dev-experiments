package main

import "fmt"

// maximumCandies возвращает максимальное число конфет, которое может получить каждый ребёнок.
func maximumCandies(candies []int, k int64) int {
	// Находим максимум в массиве candies, это будет верхней границей бинарного поиска.
	maxCandy := 0
	for _, c := range candies {
		if c > maxCandy {
			maxCandy = c
		}
	}

	low, high := 1, maxCandy
	ans := 0

	// Бинарный поиск по возможным значениям X (количество конфет на ребёнка)
	for low <= high {
		mid := (low + high) / 2
		var count int64 = 0

		// Подсчитываем, сколько кучек размера mid можно получить из всех кучек
		for _, c := range candies {
			count += int64(c / mid)
		}

		// Если можно сформировать не меньше k кучек, то значение mid допустимо
		if count >= k {
			ans = mid     // обновляем ответ
			low = mid + 1 // пробуем увеличить X
		} else {
			high = mid - 1 // уменьшаем X, так как его слишком много
		}
	}
	return ans
}

func main() {
	// Пример 1:
	candies1 := []int{5, 8, 6}
	k1 := int64(3)
	fmt.Println(maximumCandies(candies1, k1)) // Ожидаемый вывод: 5

	// Пример 2:
	candies2 := []int{2, 5}
	k2 := int64(11)
	fmt.Println(maximumCandies(candies2, k2)) // Ожидаемый вывод: 0
}
