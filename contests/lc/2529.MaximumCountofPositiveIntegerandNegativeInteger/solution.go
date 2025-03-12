package main

import (
	"fmt"
	"sort"
)

func maximumCount(nums []int) int {
	n := len(nums)

	// Count negatives: Find the first index where nums[i] is >= 0.
	negCount := sort.Search(n, func(i int) bool {
		return nums[i] >= 0
	})

	// Count positives: Find the first index where nums[i] is greater than 0.
	posIndex := sort.Search(n, func(i int) bool {
		return nums[i] > 0
	})
	posCount := n - posIndex

	if negCount > posCount {
		return negCount
	}
	return posCount
}

func main() {
	nums1 := []int{-2, -1, -1, 1, 2, 3}
	fmt.Println("Maximum count:", maximumCount(nums1)) // Expected output: 3

	nums2 := []int{-3, -2, -1, 0, 0, 1, 2}
	fmt.Println("Maximum count:", maximumCount(nums2)) // Expected output: 3

	nums3 := []int{5, 20, 66, 1314}
	fmt.Println("Maximum count:", maximumCount(nums3)) // Expected output: 4
}
