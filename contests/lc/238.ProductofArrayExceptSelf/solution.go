package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)

	// Step 1: calculate prefix products
	prefix := 1
	for i := 0; i < n; i++ {
		answer[i] = prefix
		prefix *= nums[i]
	}

	// Step 2: calculate suffix products and multiply with prefix
	suffix := 1
	for i := n - 1; i >= 0; i-- {
		answer[i] *= suffix
		suffix *= nums[i]
	}

	return answer
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))      // Output: [24 12 8 6]
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3})) // Output: [0 0 9 0 0]
}
