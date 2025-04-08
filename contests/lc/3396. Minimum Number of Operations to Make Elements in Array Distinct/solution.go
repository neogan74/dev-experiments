package main

import (
	"fmt"
	"math"
)

func minimumOperations(nums []int) int {
	m := make(map[int]bool)
	duplicateIdx := -1

	// find duplicate element from right to left
	for i := len(nums) - 1; i >= 0; i-- {
		if _, ok := m[nums[i]]; ok {
			duplicateIdx = i
			break
		} else {
			m[nums[i]] = true
		}
	}

	if duplicateIdx == -1 {
		return 0
	}

	// round up, ex: 1.1 -> 2
	return int(math.Ceil(float64(duplicateIdx+1) / 3))
}

func main() {
	// Пример 1:
	nums1 := []int{1, 2, 2}
	// Ожидаемое преобразование: [1, 2, 3], количество операций = 1.
	fmt.Printf("Пример 1:\nИсходный массив: %v\nОпераций: %d\n\n", nums1, minimumOperations(append([]int(nil), nums1...)))

	// Пример 2:
	nums2 := []int{3, 2, 1, 2, 1, 7}
	// Примерное преобразование: [1, 2, 3, 4, 5, 7] или иной вариант, итоговое количество операций = 6.
	fmt.Printf("Пример 2:\nИсходный массив: %v\nОпераций: %d\n\n", nums2, minimumOperations(append([]int(nil), nums2...)))

	// Пример 3:
	nums3 := []int{1, 1, 1, 1}
	// Примерное преобразование: [1, 2, 3, 4], количество операций = 6.
	fmt.Printf("Пример 3:\nИсходный массив: %v\nОпераций: %d\n", nums3, minimumOperations(append([]int(nil), nums3...)))
}
