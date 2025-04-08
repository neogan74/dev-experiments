package main

import (
	"fmt"
)

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	// Если сумма нечетная, деление на две равные части невозможно.
	if sum%2 != 0 {
		return false
	}

	target := sum / 2
	// dp[i] будет истинным, если можно получить сумму i из подмножества элементов.
	dp := make([]bool, target+1)
	dp[0] = true

	// Для каждого числа обновляем dp от target до num.
	for _, num := range nums {
		// Обходим в обратном порядке, чтобы не использовать один элемент несколько раз.
		for j := target; j >= num; j-- {
			// Если уже можно получить сумму (j - num), то добавив текущий элемент можно получить j.
			dp[j] = dp[j] || dp[j-num]
		}
	}

	return dp[target]
}

func main() {
	// Пример 1:
	nums1 := []int{1, 5, 11, 5}
	// Сумма = 22, target = 11. Можно разбить: [1,5,5] и [11].
	fmt.Printf("nums = %v → %v\n", nums1, canPartition(nums1)) // Ожидается true

	// Пример 2:
	nums2 := []int{1, 2, 3, 5}
	// Сумма = 11 (нечетная), поэтому нельзя разбить.
	fmt.Printf("nums = %v → %v\n", nums2, canPartition(nums2)) // Ожидается false
}
