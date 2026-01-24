package minimizemaximumpairsuminarray

import "sort"

func minPairSum(nums []int) int {
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	best := 0
	for left < right {
		sum := nums[left] + nums[right]
		if sum > best {
			best = sum
		}
		left++
		right--
	}
	return best
}

func minPairSum2(nums []int) int {
	sort.Ints(nums)
	max_pair_sum := 0
	for i := 0; i < len(nums)/2; i++ {
		max_pair_sum = max(max_pair_sum, (nums[i] + nums[len(nums)-1-i]))
	}

	return max_pair_sum
}
