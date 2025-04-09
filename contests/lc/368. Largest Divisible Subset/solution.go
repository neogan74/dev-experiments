package main

import (
	"fmt"
	"sort"
)

// largestDivisibleSubset возвращает наибольшее делимое подмножество чисел.
func largestDivisibleSubset(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return nil
	}

	// Сортируем массив.
	sort.Ints(nums)

	// dp[i] — размер наибольшего делимого подмножества, заканчивающегося nums[i]
	dp := make([]int, n)
	// prev[i] — индекс предыдущего элемента в оптимальной цепочке, заканчивающейся в i.
	prev := make([]int, n)

	// Изначально каждое подмножество состоит из одного элемента.
	for i := 0; i < n; i++ {
		dp[i] = 1
		prev[i] = -1
	}

	maxIdx := 0 // индекс, где достигается максимум dp
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			// Если nums[i] делится на nums[j] и можем увеличить длину подмножества
			if nums[i]%nums[j] == 0 && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
				prev[i] = j
			}
		}
		if dp[i] > dp[maxIdx] {
			maxIdx = i
		}
	}

	// Восстанавливаем ответ, начиная с maxIdx
	var subset []int
	for maxIdx != -1 {
		subset = append(subset, nums[maxIdx])
		maxIdx = prev[maxIdx]
	}
	// Сейчас подмножество восстановлено в обратном порядке, его можно перевернуть.
	reverse(subset)
	return subset
}

// reverse переворачивает срез чисел in-place.
func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func main() {
	// Пример 1:
	nums1 := []int{1, 2, 3}
	// Возможные варианты: [1,2] или [1,3] (так как 1 делит 2 и 1 делит 3). Наибольшего по размеру подмножества длины 2.
	fmt.Printf("nums = %v → %v\n", nums1, largestDivisibleSubset(nums1))

	// Пример 2:
	nums2 := []int{1, 2, 4, 8}
	// Явное решение: [1,2,4,8]
	fmt.Printf("nums = %v → %v\n", nums2, largestDivisibleSubset(nums2))

	// Пример 3:
	nums3 := []int{1, 3, 6, 24}
	// Возможное решение: [1,3,6,24]
	fmt.Printf("nums = %v → %v\n", nums3, largestDivisibleSubset(nums3))
}
