package main

import (
	"fmt"
	"math"
)

// minTime возвращает минимальное время, необходимое для починки всех машин.
func minTime(ranks []int, cars int) int64 {
	// Находим минимальный ранг для установки верхней границы бинарного поиска.
	minRank := ranks[0]
	for _, r := range ranks {
		if r < minRank {
			minRank = r
		}
	}

	// Верхняя граница: в худшем случае один механик с наименьшим рангом чинит все машины.
	// Время = minRank * cars^2.
	low := int64(0)
	high := int64(minRank) * int64(cars) * int64(cars)
	answer := high

	// Функция проверки: может ли время T обеспечить починку хотя бы cars машин.
	canFix := func(T int64) bool {
		var total int64 = 0
		for _, r := range ranks {
			// Максимальное количество машин, которое может починить механик с рангом r за T минут:
			// n = floor(sqrt(T / r))
			total += int64(math.Floor(math.Sqrt(float64(T) / float64(r))))
			// Если уже набрали достаточно, можно выйти раньше.
			if total >= int64(cars) {
				return true
			}
		}
		return total >= int64(cars)
	}

	// Бинарный поиск по времени.
	for low <= high {
		mid := (low + high) / 2
		if canFix(mid) {
			answer = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return answer
}

func main() {
	// Пример 1:
	ranks1 := []int{4, 2, 3, 1}
	cars1 := 10
	fmt.Println("func1", minTime(ranks1, cars1)) // Ожидаемый вывод: 16
	fmt.Println("func2", repairCars2(ranks1, cars1))

	// Пример 2:
	ranks2 := []int{5, 1, 8}
	cars2 := 6
	fmt.Println("func1", minTime(ranks2, cars2)) // Ожидаемый вывод: 16
	fmt.Println("func2", repairCars2(ranks2, cars2))
}
