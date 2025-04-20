package _563_CountTheNumberOfFairPairs

import (
	"sort"
)

// countFairPairs возвращает количество пар (i, j), i<j, таких что
// lower <= nums[i] + nums[j] <= upper.
func countFairPairs(nums []int, lower, upper int) int64 {
	sort.Ints(nums)
	n := len(nums)
	var count int64

	for i := 0; i < n-1; i++ {
		// Ищем в срезе nums[i+1:] индексы первой позиции >= lower-nums[i]
		// и первой позиции > upper-nums[i].
		lo := lower - nums[i]
		hi := upper - nums[i]

		// длина оставшейся части массива
		m := n - (i + 1)
		// индекс в nums[i+1:] первого элемента >= lo
		lowIdx := sort.Search(m, func(k int) bool {
			return nums[i+1+k] >= lo
		})
		// индекс в nums[i+1:] первого элемента > hi
		highIdx := sort.Search(m, func(k int) bool {
			return nums[i+1+k] > hi
		})

		// все индексы [lowIdx, highIdx) дают суммы в [lower, upper]
		count += int64(highIdx - lowIdx)
	}
	return count
}
