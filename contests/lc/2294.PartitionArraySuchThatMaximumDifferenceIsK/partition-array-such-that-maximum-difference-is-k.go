package _2294_partition_array_such_that_maximum_difference_is_k

import "sort"

func partitionArray(nums []int, k int) int {
	sort.Ints(nums)
	ans, a := 1, nums[0]
	for _, b := range nums {
		if b-a > k {
			a = b
			ans++
		}
	}
	return ans
}
