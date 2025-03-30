package main

import "fmt"

// jump возвращает минимальное количество прыжков, необходимых для достижения конца массива.
func jump(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	jumps, currentEnd, furthest := 0, 0, 0

	// Проходим по массиву, кроме последнего элемента.
	for i := 0; i < n-1; i++ {
		// Обновляем максимальную доступную позицию
		if i+nums[i] > furthest {
			furthest = i + nums[i]
		}
		// Если дошли до конца текущего диапазона, увеличиваем количество прыжков
		if i == currentEnd {
			jumps++
			currentEnd = furthest
		}
	}
	return jumps
}

func main() {
	// Пример 1:
	nums1 := []int{2, 3, 1, 1, 4}
	fmt.Println(jump(nums1)) // Ожидаемый вывод: 2

	// Пример 2:
	nums2 := []int{2, 3, 0, 1, 4}
	fmt.Println(jump(nums2)) // Ожидаемый вывод: 2
}
