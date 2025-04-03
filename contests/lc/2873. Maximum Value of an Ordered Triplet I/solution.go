package main

import (
	"fmt"
)

func maximumTripletValue(nums []int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}

	// leftMax[j] будет хранить максимум среди элементов с индексами от 0 до j-1.
	leftMax := make([]int, n)
	leftMax[0] = nums[0]
	for j := 1; j < n; j++ {
		if nums[j-1] > leftMax[j-1] {
			leftMax[j] = nums[j-1]
		} else {
			leftMax[j] = leftMax[j-1]
		}
	}

	// suffix[j] будет хранить максимум среди элементов с индексами от j до n-1.
	suffix := make([]int, n)
	suffix[n-1] = nums[n-1]
	for j := n - 2; j >= 0; j-- {
		if nums[j] > suffix[j+1] {
			suffix[j] = nums[j]
		} else {
			suffix[j] = suffix[j+1]
		}
	}

	maxValue := 0
	// Перебираем все возможные j, где j от 1 до n-2
	for j := 1; j <= n-2; j++ {
		diff := leftMax[j] - nums[j]
		// Если разность отрицательная, то (nums[i] - nums[j]) < 0, и произведение с положительным nums[k] будет отрицательным.
		if diff < 0 {
			continue
		}
		// suffix[j+1] содержит максимум для индексов от j+1 до n-1 (то есть оптимальный выбор для k)
		candidate := diff * suffix[j+1]
		if candidate > maxValue {
			maxValue = candidate
		}
	}
	return maxValue
}

func main() {
	// Пример 1:
	nums1 := []int{12, 6, 1, 2, 7}
	// Ожидаемый ответ: 77, например, для тройки (0,2,4): (12-1)*7 = 77.
	fmt.Printf("Пример 1: %v → %d\n", nums1, maximumTripletValue(nums1))

	// Пример 2:
	nums2 := []int{1, 10, 3, 4, 19}
	// Ожидаемый ответ: 133, например, для тройки (1,2,4): (10-3)*19 = 133.
	fmt.Printf("Пример 2: %v → %d\n", nums2, maximumTripletValue(nums2))

	// Пример 3:
	nums3 := []int{1, 2, 3}
	// Единственная тройка (0,1,2) даёт (1-2)*3 = -3, отрицательное, поэтому ответ 0.
	fmt.Printf("Пример 3: %v → %d\n", nums3, maximumTripletValue(nums3))
}
