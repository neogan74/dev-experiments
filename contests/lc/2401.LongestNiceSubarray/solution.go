package main

import (
	"fmt"
)

func longestNiceSubarray(nums []int) int {
	maxLength := 0
	bits, left := 0, 0

	for right := 0; right < len(nums); right++ {
		// Shrink window until nums[right] doesn't overlap with bits
		for (bits & nums[right]) != 0 {
			bits ^= nums[left] // remove nums[left] from current bits
			left++
		}

		bits |= nums[right] // include nums[right] bits
		maxLength = max(maxLength, right-left+1)
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestNiceSubarray([]int{1, 3, 8, 48, 10})) // Output: 3
	fmt.Println(longestNiceSubarray([]int{3, 1, 5, 11, 13})) // Output: 1
}
