package main

import (
	"fmt"
	"sort"
)

func minCapability(nums []int, k int) int {
	return sort.Search(1e9, func(mx int) bool {
		cnt, n := 0, len(nums)
		for i := 0; i < n; i++ {
			if nums[i] <= mx {
				cnt++
				i++
			}
		}
		return cnt >= k
	})
}

func main() {
	// Case1
	nums := []int{2, 3, 5, 9}
	k := 2
	res := minCapability(nums, k)
	fmt.Printf("Result case1: %v\n", res)

	//Case2
	nums2 := []int{2, 7, 9, 3, 1}
	k = 2
	res2 := minCapability(nums2, k)
	fmt.Printf("Result case1: %v\n", res2)
}
