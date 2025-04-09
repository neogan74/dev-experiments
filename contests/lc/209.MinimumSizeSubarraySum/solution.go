package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	minLen := math.MaxInt32
	sum, left := 0, 0

	for right := 0; right < len(nums); right++ {
		sum += nums[right]

		for sum >= target {
			minLen = min(minLen, right-left+1)
			sum -= nums[left]
			left++
		}
	}

	if minLen == math.MaxInt32 {
		return 0
	}
	return minLen
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))        // 2
	fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))                 // 1
	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1})) // 0
}
