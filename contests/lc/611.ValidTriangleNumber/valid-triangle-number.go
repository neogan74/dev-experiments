package validtrianglenumber

import "sort"

// triangleNumber counts index triplets (i, j, k) whose values can form
// a non-degenerate triangle. After sorting we iterate the first two sides
// and binary search the longest valid third side boundary in the suffix.
func triangleNumber(nums []int) int {
	sort.Ints(nums)
	ans := 0
	for i, n := 0, len(nums); i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			left, right := j+1, n
			for left < right {
				mid := (left + right) >> 1
				if nums[mid] >= nums[i]+nums[j] {
					right = mid
				} else {
					left = mid + 1
				}
			}
			ans += left - j - 1
		}
	}
	return ans
}
