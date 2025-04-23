package main

import (
	"fmt"
)

// zeroArrayTransformation returns the minimum k such that after processing the first k queries,
// it is possible to decrement each element of nums to 0 (i.e. available decrement >= nums[j] for all j).
// If no such k exists, it returns -1.
func zeroArrayTransformation(nums []int, queries [][]int) int {
	n := len(nums)
	m := len(queries)

	// Pre-check: if nums is already a zero array, return 0.
	allZero := true
	for _, v := range nums {
		if v != 0 {
			allZero = false
			break
		}
	}
	if allZero {
		return 0
	}

	// Helper function to check if using the first k queries is sufficient.
	canTransform := func(k int) bool {
		// Create a difference array of length n+1 (extra element to handle r+1 subtraction)
		diff := make([]int, n+1)
		// Apply the first k queries to the difference array.
		for i := 0; i < k; i++ {
			l, r, val := queries[i][0], queries[i][1], queries[i][2]
			diff[l] += val
			if r+1 < n {
				diff[r+1] -= val
			}
		}
		// Build the available decrement array by computing the cumulative sum of diff.
		available := 0
		for j := 0; j < n; j++ {
			available += diff[j]
			// For each index j, available decrement must be at least nums[j].
			if available < nums[j] {
				return false
			}
		}
		return true
	}

	// Binary search for the smallest k.
	low, high := 1, m
	answer := -1
	for low <= high {
		mid := (low + high) / 2
		if canTransform(mid) {
			answer = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return answer
}

// Greedily increment k while going through each number

func minZeroArray2(nums []int, queries [][]int) int {
	increments := make([]int, len(nums)+1)
	k := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		for sum+increments[i] < nums[i] {
			if k >= len(queries) { // running out of queries
				return -1
			}
			query := queries[k]
			l := query[0]
			r := query[1]
			val := query[2]
			// if r < i, the query doesn't improve increments[i], so just keep going to the next query
			if r >= i {
				// if l < i, we mark increments[i] so that it's effective immediately
				increments[max(l, i)] += val
				// subtract it down when getting outside of the range
				increments[r+1] -= val
			}
			k++
		}
		sum += increments[i]
	}
	return k
}

func main() {
	// Example 1:
	nums1 := []int{2, 0, 2}
	queries1 := [][]int{{0, 2, 1}, {0, 2, 1}, {1, 1, 3}}
	fmt.Println(zeroArrayTransformation(nums1, queries1)) // Expected output: 2

	// Example 2:
	nums2 := []int{4, 3, 2, 1}
	queries2 := [][]int{{1, 3, 2}, {0, 2, 1}}
	fmt.Println(zeroArrayTransformation(nums2, queries2)) // Expected output: -1

	// Additional test case:
	nums3 := []int{0}
	queries3 := [][]int{{0, 0, 2}, {0, 0, 4}, {0, 0, 4}, {0, 0, 3}, {0, 0, 5}}
	fmt.Println(zeroArrayTransformation(nums3, queries3)) // Expected output: 0

	// Example 1:
	nums4 := []int{2, 0, 2}
	queries4 := [][]int{{0, 2, 1}, {0, 2, 1}, {1, 1, 3}}
	fmt.Println(minZeroArray2(nums4, queries4)) // Expected output: 2

	// Example 2:
	nums5 := []int{4, 3, 2, 1}
	queries5 := [][]int{{1, 3, 2}, {0, 2, 1}}
	fmt.Println(minZeroArray2(nums5, queries5)) // Expected output: -1

	// Additional test case:
	nums6 := []int{0}
	queries6 := [][]int{{0, 0, 2}, {0, 0, 4}, {0, 0, 4}, {0, 0, 3}, {0, 0, 5}}
	fmt.Println(minZeroArray2(nums6, queries6)) // Expected output: 0

}
