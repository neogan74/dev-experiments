package main

import (
	"fmt"
)

func maximumTripletValue(nums []int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}

	// leftMax[j] хранит максимум среди элементов с индексами от 0 до j-1.
	leftMax := make([]int, n)
	leftMax[0] = nums[0]
	for j := 1; j < n; j++ {
		// Выбираем максимум между предыдущим leftMax и nums[j-1]
		if nums[j-1] > leftMax[j-1] {
			leftMax[j] = nums[j-1]
		} else {
			leftMax[j] = leftMax[j-1]
		}
	}

	// suffixMax[i] хранит максимум среди элементов с индексами от i до n-1.
	rightMax := make([]int, n)
	rightMax[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		if nums[i] > rightMax[i+1] {
			rightMax[i] = nums[i]
		} else {
			rightMax[i] = rightMax[i+1]
		}
	}

	maxValue := 0
	// Перебираем возможные j от 1 до n-2
	for j := 1; j <= n-2; j++ {
		// Разность между максимально возможным left и текущим nums[j]
		diff := leftMax[j] - nums[j]
		if diff < 0 {
			// Если разность отрицательная, такой кандидат не даст положительного произведения
			continue
		}
		// Для k оптимально выбрать rightMax из диапазона j+1...n-1
		candidate := diff * rightMax[j+1]
		if candidate > maxValue {
			maxValue = candidate
		}
	}

	return maxValue
}

func main() {
	// Пример 1:
	nums1 := []int{12, 6, 1, 2, 7}
	// Для тройки (0, 2, 4): (12 - 1) * 7 = 77.
	fmt.Printf("Пример 1: nums = %v → %d\n", nums1, maximumTripletValue(nums1))

	// Пример 2:
	nums2 := []int{1, 10, 3, 4, 19}
	// Для тройки (1, 2, 4): (10 - 3) * 19 = 133.
	fmt.Printf("Пример 2: nums = %v → %d\n", nums2, maximumTripletValue(nums2))

	// Пример 3:
	nums3 := []int{1, 2, 3}
	// Единственная тройка (0, 1, 2) даёт (1 - 2) * 3 = -3, поэтому ответ = 0.
	fmt.Printf("Пример 3: nums = %v → %d\n", nums3, maximumTripletValue(nums3))
}
