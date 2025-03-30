package main

import "fmt"

// canJump returns true if you can reach the last index of nums, false otherwise.
func canJump(nums []int) bool {
	n := len(nums)
	// maxReach stores the farthest index that can be reached so far.
	maxReach := 0

	// Iterate over each index in the array.
	for i := 0; i < n; i++ {
		// If the current index is not reachable, then return false.
		if i > maxReach {
			return false
		}
		// Update the maximum reachable index.
		if i+nums[i] > maxReach {
			maxReach = i + nums[i]
		}
		// If we've reached or passed the last index, return true.
		if maxReach >= n-1 {
			return true
		}
	}
	return false
}

func main() {
	// Example 1:
	nums1 := []int{2, 3, 1, 1, 4}
	fmt.Println(canJump(nums1)) // Expected output: true

	// Example 2:
	nums2 := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(nums2)) // Expected output: false
}

/*
Explanation Walk‐through
Initialization:
We start with maxReach = 0.

Iteration and Updating maxReach:

At each index i, we first check if i is within the range we can reach (i <= maxReach).
If it is, we compute i + nums[i] to see how far we can jump from this position and update maxReach accordingly.
If maxReach becomes greater than or equal to the last index (n-1), we know that reaching the end is possible and return true immediately.
Failure Condition:
If we reach an index that is beyond our maxReach, it means we are stuck and cannot proceed further, so we return false.

This greedy solution runs in O(n) time and O(1) space, making it efficient for the input constraints provided.
*/
