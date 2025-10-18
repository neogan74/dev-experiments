package _397_MaximumNumberofDistinctElementsAfterOperations

import "sort"

func maxDistinctElements(nums []int, k int) int {
	sort.Ints(nums)
	ans := 0
	const NEG_INF = -1 << 60
	occupied := NEG_INF
	for _, x := range nums {
		place := occupied + 1
		left := x - k
		right := x + k
		if place < left {
			place = left
		}
		if place <= right {
			occupied = place
			ans++
		}
	}
	return ans
}
