package maximumuniquesubarraysumafterdeletion

import "math"

func maximumUniqueSubarraySumAfterDeletion(nums []int) int {
	unique := map[int]int{}
	smallestNegative := -math.MaxInt
	sum := 0
	found := false
	for _, n := range nums {
		if _, ok := unique[n]; !ok {
			if n >= 0 {
				unique[n]++
				sum += n
				found = true
			} else {
				smallestNegative = max(smallestNegative, n)
			}
		}
	}
	if !found {
		return smallestNegative
	}
	return sum
}
