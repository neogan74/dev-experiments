package main

import (
	"fmt"
)

func minOperations(nums []int, k int) int {
	uniqueValues := make(map[int]bool)
	for _, num := range nums {
		if num < k {
			return -1
		}
		if num > k {
			uniqueValues[num] = true
		}
	}
	return len(uniqueValues)
}

func main() {
	// Примеры тестов
	testCases := []struct {
		nums []int
		k    int
	}{
		{[]int{5, 2, 5, 4, 5}, 2},
		{[]int{2, 1, 2}, 2},
		{[]int{9, 7, 5, 3}, 1},
		{[]int{4, 4, 4}, 4},
		{[]int{10, 10, 9, 8, 7}, 7},
	}

	for i, tc := range testCases {
		result := minOperations(tc.nums, tc.k)
		fmt.Printf("Test case %d: nums = %v, k = %d → result = %d\n", i+1, tc.nums, tc.k, result)
	}
}
