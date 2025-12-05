package main

// countPartitions returns the number of partition indices where the difference
// between left and right subarray sums is even.
// The difference is 2*prefix - total, so its parity matches the total sum.
// If the total is even, every split (n-1 of them) works; if odd, none do.
func countPartitions(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}

	if total%2 != 0 {
		return 0
	}
	return len(nums) - 1
}
