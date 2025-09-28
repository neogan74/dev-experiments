package main

import "sort"

// largestPerimeter returns the maximum perimeter of any triangle that can be
// formed from nums; returns 0 when no valid triangle exists.
func largestPerimeter(nums []int) int {
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 2; i-- {
		if sum := nums[i-1] + nums[i-2]; sum > nums[i] {
			return sum + nums[i]
		}
	}
	return 0
}

func main() {}
